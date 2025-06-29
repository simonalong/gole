package orm

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	driverMysql "github.com/go-sql-driver/mysql"
	"github.com/lib/pq"
	"github.com/qustavo/sqlhooks/v2"
	"github.com/simonalong/gole/config"
	baseLogger "github.com/simonalong/gole/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"strings"
	"time"
)

var GlobalBaseMap map[string]*gorm.DB = make(map[string]*gorm.DB)

func NewGormClient() (*gorm.DB, error) {
	if !config.GetValueBoolDefault("gole.datasource.enable", false) {
		return nil, errors.New("数据库配置 gole.datasource.enable 为false；请开启")
	}
	return doNewGormDb("", &gorm.Config{})
}

func NewGormClientWitConfig(gormConfig *gorm.Config) (*gorm.DB, error) {
	if !config.GetValueBoolDefault("gole.datasource.enable", false) {
		return nil, errors.New("数据库配置 gole.datasource.enable 为false；请开启")
	}
	return doNewGormDb("", gormConfig)
}

func NewGormClientWithName(datasourceName string) (*gorm.DB, error) {
	if !config.GetValueBoolDefault("gole.datasource.enable", false) {
		return nil, errors.New("数据库配置 gole.datasource.enable 为false；请开启")
	}
	return doNewGormDb(datasourceName, &gorm.Config{})
}

func NewGormClientWithNameAndConfig(datasourceName string, gormConfig *gorm.Config) (*gorm.DB, error) {
	if !config.GetValueBoolDefault("gole.datasource.enable", false) {
		return nil, errors.New("数据库配置 gole.datasource.enable 为false；请开启")
	}
	return doNewGormDb(datasourceName, gormConfig)
}

func doNewGormDb(datasourceName string, gormConfig *gorm.Config) (*gorm.DB, error) {
	datasourceConfig := DatasourceConfig{}
	targetDatasourceName := "gole.datasource"
	if datasourceName != "" {
		targetDatasourceName = "gole.datasource." + datasourceName
	}
	err := config.GetValueObject(targetDatasourceName, &datasourceConfig)
	if err != nil {
		baseLogger.Warn("读取读取配置【datasource】异常")
		return nil, err
	}

	// 注册原生的sql的hook
	if len(gormHooks) != 0 {
		sqlRegister(datasourceConfig.DriverName)
	}

	var gormDb *gorm.DB
	dsn := getDbDsn(datasourceConfig.DriverName, datasourceConfig)
	gormDb, err = gorm.Open(getDialect(dsn, datasourceConfig.DriverName), gormConfig)
	if err != nil {
		baseLogger.Warnf("连接数据库db异常：%v", err.Error())
		return nil, err
	}

	_db, err := gormDb.DB()
	if err != nil {
		baseLogger.Warnf("获取数据库db异常：%v", err.Error())
		return nil, err
	}

	maxIdleConns := config.GetValueInt("gole.datasource.connect-pool.max-idle-conns")
	if maxIdleConns != 0 {
		// 设置空闲的最大连接数
		_db.SetMaxIdleConns(maxIdleConns)
	}

	maxOpenConns := config.GetValueInt("gole.datasource.connect-pool.max-open-conns")
	if maxOpenConns != 0 {
		// 设置数据库打开连接的最大数量
		_db.SetMaxOpenConns(maxOpenConns)
	}

	maxLifeTime := config.GetValueString("gole.datasource.connect-pool.max-life-time")
	if maxLifeTime != "" {
		// 设置连接可重复使用的最大时间
		t, err := time.ParseDuration(maxLifeTime)
		if err != nil {
			baseLogger.Warn("读取配置【gole.datasource.connect-pool.max-life-time】异常", err)
		} else {
			_db.SetConnMaxLifetime(t)
		}
	}

	maxIdleTime := config.GetValueString("gole.datasource.connect-pool.max-idle-time")
	if maxIdleTime != "" {
		// 设置conn最大空闲时间设置连接空闲的最大时间
		t, err := time.ParseDuration(maxIdleTime)
		if err != nil {
			baseLogger.Warn("读取配置【gole.datasource.connect-pool.max-idle-time】异常", err)
		} else {
			_db.SetConnMaxIdleTime(t)
		}
	}

	gormDb.Logger = &GormLoggerAdapter{}

	//pGormDb := &GormClient{GormDb: gormDb}
	GlobalBaseMap[datasourceName] = gormDb

	return gormDb, nil
}

// 特殊字符处理
func specialCharChange(url string) string {
	return strings.ReplaceAll(url, "/", "%2F")
}

func getDialect(dsn, driverName string) gorm.Dialector {
	switch driverName {
	case "mysql":
		return mysql.New(getMysqlConfig(dsn, driverName))
	case "postgresql":
		return postgres.New(postgres.Config{DSN: dsn, DriverName: wrapDriverName(driverName)})
	case "sqlserver":
		return sqlserver.New(sqlserver.Config{DSN: dsn, DriverName: wrapDriverName(driverName)})
	}
	return nil
}

func sqlRegister(driverName string) {
	name := wrapDriverName(driverName)
	for _, driver := range sql.Drivers() {
		if driver == name {
			return
		}
	}

	switch driverName {
	case "mysql":
		sql.Register(name, sqlhooks.Wrap(&driverMysql.MySQLDriver{}, &BaseSqlHookProxy{DriverName: driverName}))
	case "postgresql":
		sql.Register(name, sqlhooks.Wrap(&pq.Driver{}, &BaseSqlHookProxy{DriverName: driverName}))
	}
}

func getMysqlConfig(dsn, driverName string) mysql.Config {
	return mysql.Config{
		DriverName:                    driverName,
		DSN:                           dsn,
		ServerVersion:                 config.GetValueStringDefault("gole.datasource.mysql.server-version", ""),
		SkipInitializeWithVersion:     config.GetValueBoolDefault("gole.datasource.mysql.skip-initialize-with-version", false),
		DefaultStringSize:             config.GetValueUIntDefault("gole.datasource.mysql.default-string-size", 0),
		DisableWithReturning:          config.GetValueBoolDefault("gole.datasource.mysql.disable-with-returning", false),
		DisableDatetimePrecision:      config.GetValueBoolDefault("gole.datasource.mysql.disable-datetime-precision", false),
		DontSupportRenameIndex:        config.GetValueBoolDefault("gole.datasource.mysql.dont-support-rename-index", false),
		DontSupportRenameColumn:       config.GetValueBoolDefault("gole.datasource.mysql.dont-support-rename-column", false),
		DontSupportForShareClause:     config.GetValueBoolDefault("gole.datasource.mysql.dont-support-for-share-clause", false),
		DontSupportNullAsDefaultValue: config.GetValueBoolDefault("gole.datasource.mysql.dont-support-null-as-default-value", false),
	}
}

func wrapDriverName(driverName string) string {
	if len(gormHooks) != 0 {
		return driverName + "Hook"
	}
	return driverName
}

type BaseGormHook interface {
	Before(ctx context.Context, driverName string, parameters map[string]any) (context.Context, error)
	After(ctx context.Context, driverName string, parameters map[string]any) (context.Context, error)
	Err(ctx context.Context, driverName string, err error, parameters map[string]any) error
}

var gormHooks []BaseGormHook

func init() {
	gormHooks = []BaseGormHook{}
}

func AddGormHook(hook BaseGormHook) {
	if !config.GetValueBoolDefault("gole.meter.orm.enable", true) {
		return
	}
	gormHooks = append(gormHooks, hook)
}

type BaseSqlHookProxy struct {
	DriverName string
}

func (proxy *BaseSqlHookProxy) Before(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	var ctxFinal context.Context
	for _, hook := range gormHooks {
		parametersMap := map[string]any{
			"query": query,
			"args":  args,
		}
		_ctx, err := hook.Before(ctx, proxy.DriverName, parametersMap)
		if err != nil {
			return _ctx, err
		} else {
			ctxFinal = _ctx
		}
	}
	return ctxFinal, nil
}

func (proxy *BaseSqlHookProxy) After(ctx context.Context, query string, args ...interface{}) (context.Context, error) {
	for _, hook := range gormHooks {
		parametersMap := map[string]any{
			"query": query,
			"args":  args,
		}
		ctx, err := hook.After(ctx, proxy.DriverName, parametersMap)
		if err != nil {
			return ctx, err
		}
	}
	return ctx, nil
}

func (proxy *BaseSqlHookProxy) OnError(ctx context.Context, err error, query string, args ...interface{}) error {
	for _, hook := range gormHooks {
		parametersMap := map[string]any{
			"query": query,
			"args":  args,
		}
		err := hook.Err(ctx, proxy.DriverName, err, parametersMap)
		if err != nil {
			return err
		}
	}
	return nil
}

type GormLoggerAdapter struct {
}

func (l *GormLoggerAdapter) LogMode(level logger.LogLevel) logger.Interface {
	var levelStr logrus.Level
	switch level {
	case logger.Silent:
		levelStr = logrus.TraceLevel
	case logger.Error:
		levelStr = logrus.ErrorLevel
	case logger.Warn:
		levelStr = logrus.WarnLevel
	case logger.Info:
		levelStr = logrus.InfoLevel
	}
	baseLogger.Group("orm").SetLevel(levelStr)
	return l
}

func (l *GormLoggerAdapter) Info(ctx context.Context, msg string, data ...interface{}) {
	baseLogger.Infof(msg, data)
}

func (l *GormLoggerAdapter) Warn(ctx context.Context, msg string, data ...interface{}) {
	baseLogger.Warnf(msg, data)
}

func (l *GormLoggerAdapter) Error(ctx context.Context, msg string, data ...interface{}) {
	baseLogger.Errorf(msg, data)
}

func (l *GormLoggerAdapter) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sqlStr, rowsAffected := fc()
	if err != nil {
		if err.Error() == "record not found" {
			baseLogger.Group("orm").Debugf("[SQL][%v]%s; error: %v", elapsed, sqlStr, err.Error())
			return
		}
		baseLogger.Group("orm").Errorf("[SQL][%v]%s; error: %v", elapsed, sqlStr, err.Error())
	} else {
		baseLogger.Group("orm").Debugf("[SQL][%v][row:%v]%s", elapsed, rowsAffected, sqlStr)
	}
}

func GlobalGormContextLoad() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, gormDb := range GlobalBaseMap {
			gormDb.WithContext(c.Request.Context())
		}
	}
}
