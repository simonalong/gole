package orm

import (
	"context"
	"fmt"
	"gitlab.seatakcloud.com/cbb/base/cbb-base-boot/constants"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/bean"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/config"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/global"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/listener"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/logger"
	baseTime "gitlab.seatakcloud.com/cbb/base/cbb-base/time"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/util"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	"go.opentelemetry.io/otel/trace"
	"reflect"
	"strings"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/contexts"
	"xorm.io/xorm/log"
)

type BaseXormHook interface {
	BeforeProcess(c *contexts.ContextHook, driverName string) (context.Context, error)
	AfterProcess(c *contexts.ContextHook, driverName string) error
}

var defaultXormHooks []DefaultXormHook

type DefaultXormHook struct {
	driverName   string
	baseXormHook BaseXormHook
}

func (defaultHook *DefaultXormHook) BeforeProcess(c *contexts.ContextHook) (context.Context, error) {
	return defaultHook.baseXormHook.BeforeProcess(c, defaultHook.driverName)
}

func (defaultHook *DefaultXormHook) AfterProcess(c *contexts.ContextHook) error {
	return defaultHook.baseXormHook.AfterProcess(c, defaultHook.driverName)
}

type OtelXormHook struct {
	datasourceName string
	driverName     string
	tracer         trace.Tracer
	attrs          []attribute.KeyValue
}

func (otelXormHook *OtelXormHook) BeforeProcess(contextHook *contexts.ContextHook) (context.Context, error) {
	spanName := contextHook.SQL
	if contextHook.SQL != "" {
		spanName = strings.SplitN(contextHook.SQL, " ", 2)[0]
	}
	ctx, _ := otelXormHook.tracer.Start(global.GetGlobalContext(), "GormClient: "+spanName, trace.WithSpanKind(trace.SpanKindClient))
	return ctx, nil
}

func (otelXormHook *OtelXormHook) AfterProcess(contextHook *contexts.ContextHook) error {
	span := trace.SpanFromContext(contextHook.Ctx)
	defer span.End()

	attrs := make([]attribute.KeyValue, 0, len(otelXormHook.attrs)+4)
	attrs = append(attrs, otelXormHook.attrs...)

	if sys := dbSystem(otelXormHook.driverName); sys.Valid() {
		attrs = append(attrs, sys)
	}

	attrs = append(attrs, semconv.DBQueryTextKey.String(contextHook.SQL))
	attrs = append(attrs, attribute.Key("GormClient.datasource.name").String(otelXormHook.datasourceName))
	attrs = append(attrs, attribute.Key("execute.time").String(baseTime.ParseDurationForView(contextHook.ExecuteTime)))
	var argStrSlice []string
	for _, arg := range contextHook.Args {
		if reflect.TypeOf(arg) == reflect.TypeOf(time.Time{}) {
			argStrSlice = append(argStrSlice, baseTime.TimeToStringYmdHmsS(arg.(time.Time)))
		} else {
			argStrSlice = append(argStrSlice, util.ToString(arg))
		}
	}
	attrs = append(attrs, attribute.Key("sql.args").StringSlice(argStrSlice))

	if contextHook.Result != nil {
		rowsAffected, err := contextHook.Result.RowsAffected()
		if rowsAffected != -1 {
			attrs = append(attrs, attribute.Key("rows.affected").Int64(rowsAffected))
		}
		if err != nil {
			attrs = append(attrs, attribute.Key("GormClient.err").String(err.Error()))
			span.SetAttributes(attrs...)

			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
		}
		return err
	}

	if contextHook.Err != nil {
		attrs = append(attrs, attribute.Key("GormClient.err").String(contextHook.Err.Error()))
		span.SetAttributes(attrs...)

		span.RecordError(contextHook.Err)
		span.SetStatus(codes.Error, contextHook.Err.Error())
	} else {
		span.SetAttributes(attrs...)
	}
	return contextHook.Err
}

func dbSystem(driverName string) attribute.KeyValue {
	switch driverName {
	case "mysql":
		return semconv.DBSystemMySQL
	case "mssql":
		return semconv.DBSystemMSSQL
	case "postgres", "postgresql":
		return semconv.DBSystemPostgreSQL
	case "sqlserver":
		return semconv.DBSystemKey.String("sqlserver")
	case "clickhouse":
		return semconv.DBSystemKey.String("clickhouse")
	default:
		return attribute.KeyValue{}
	}
}

func init() {
	defaultXormHooks = []DefaultXormHook{}
}

func NewXormClient() (*xorm.Engine, error) {
	return doNewXormDb("", map[string]string{})
}

func NewXormDbWithParams(params map[string]string) (*xorm.Engine, error) {
	return doNewXormDb("", params)
}

func NewXormDbWithName(datasourceName string) (*xorm.Engine, error) {
	return doNewXormDb(datasourceName, map[string]string{})
}

func NewXormDbWithNameParams(datasourceName string, params map[string]string) (*xorm.Engine, error) {
	return doNewXormDb(datasourceName, params)
}

func AddXormHook(hook BaseXormHook) {
	defaultXormHook := DefaultXormHook{baseXormHook: hook}
	defaultXormHooks = append(defaultXormHooks, defaultXormHook)
	xormDbs := bean.GetBeanWithNamePre(constants.BeanNameXormPre)
	if xormDbs == nil {
		return
	}
	for _, db := range xormDbs {
		db.(*xorm.Engine).AddHook(&defaultXormHook)
	}
}

func doNewXormDb(datasourceName string, params map[string]string) (*xorm.Engine, error) {
	datasourceConfig := DatasourceConfig{}
	targetDatasourceName := "base.datasource"
	if datasourceName != "" {
		targetDatasourceName = "base.datasource." + datasourceName
	}
	err := config.GetValueObject(targetDatasourceName, &datasourceConfig)
	if err != nil {
		logger.Warn("读取读取配置【datasource】异常")
		return nil, err
	}

	var dsn = getDbDsn(datasourceConfig.DriverName, datasourceConfig)
	var xormDb *xorm.Engine
	xormDb, err = xorm.NewEngineWithParams(datasourceConfig.DriverName, dsn, params)
	if err != nil {
		logger.Warnf("获取数据库db异常：%v", err.Error())
		return nil, err
	}

	for _, hook := range defaultXormHooks {
		hook.driverName = datasourceConfig.DriverName
		xormDb.AddHook(&hook)
	}

	maxIdleConns := config.GetValueInt("base.datasource.connect-pool.max-idle-conns")
	if maxIdleConns != 0 {
		// 设置空闲的最大连接数
		xormDb.SetMaxIdleConns(maxIdleConns)
	}

	maxOpenConns := config.GetValueInt("base.datasource.connect-pool.max-open-conns")
	if maxOpenConns != 0 {
		// 设置数据库打开连接的最大数量
		xormDb.SetMaxOpenConns(maxOpenConns)
	}

	maxLifeTime := config.GetValueString("base.datasource.connect-pool.max-life-time")
	if maxLifeTime != "" {
		// 设置连接可重复使用的最大时间
		t, err := time.ParseDuration(maxLifeTime)
		if err != nil {
			logger.Warn("读取配置【base.datasource.connect-pool.max-life-time】异常", err)
		} else {
			xormDb.SetConnMaxLifetime(t)
		}
	}

	xormDb.ShowSQL(true)
	xormDb.SetLogger(&XormLoggerAdapter{})
	bean.AddBean(constants.BeanNameXormPre+datasourceName, xormDb)

	// 支持opentelemetry埋点
	if config.GetValueBoolDefault("base.opentelemetry.enable", false) {
		xormDb.AddHook(&OtelXormHook{
			datasourceName: datasourceName,
			driverName:     datasourceConfig.DriverName,
			tracer:         global.Tracer,
		})
	}

	// 添加orm的配置监听器
	listener.AddListener(listener.EventOfConfigChange, ConfigChangeListenerOfOrm)
	return xormDb, nil
}

func NewXormDbMasterSlave(masterDatasourceName string, slaveDatasourceNames []string, policies ...xorm.GroupPolicy) (*xorm.EngineGroup, error) {
	masterDb, err := NewXormDbWithName(masterDatasourceName)
	if err != nil {
		logger.Warnf("获取数据库 主节点【%v】失败，%v", masterDatasourceName, err.Error())
		return nil, err
	}

	var slaveDbs []*xorm.Engine
	for _, slaveDatasource := range slaveDatasourceNames {
		slaveDb, err := NewXormDbWithName(slaveDatasource)
		if err != nil {
			logger.Warnf("获取数据库 从节点【%v】失败，%v", slaveDatasource, err.Error())
			return nil, err
		}

		slaveDbs = append(slaveDbs, slaveDb)
	}

	return xorm.NewEngineGroup(masterDb, slaveDbs, policies...)
}

// LoggerAdapter wraps a Logger interface as LoggerContext interface
type XormLoggerAdapter struct {
}

// BeforeSQL implements ContextLogger
func (l *XormLoggerAdapter) BeforeSQL(ctx log.LogContext) {}

// AfterSQL implements ContextLogger
func (l *XormLoggerAdapter) AfterSQL(ctx log.LogContext) {
	var sessionPart string
	v := ctx.Ctx.Value("__xorm_session_id")
	if key, ok := v.(string); ok {
		sessionPart = fmt.Sprintf(" [%s]", key)
	}
	if ctx.ExecuteTime > 0 {
		logger.Group("orm").Debugf("[SQL]%s %s %v - %v", sessionPart, ctx.SQL, ctx.Args, ctx.ExecuteTime)
	} else {
		logger.Group("orm").Debugf("[SQL]%s %s %v", sessionPart, ctx.SQL, ctx.Args)
	}
}

// Debugf implements ContextLogger
func (l *XormLoggerAdapter) Debugf(format string, v ...interface{}) {
	logger.Group("orm").Debug(format, v)
}

// Errorf implements ContextLogger
func (l *XormLoggerAdapter) Errorf(format string, v ...interface{}) {
	logger.Errorf(format, v)
}

// Infof implements ContextLogger
func (l *XormLoggerAdapter) Infof(format string, v ...interface{}) {
	logger.Infof(format, v)
}

// Warnf implements ContextLogger
func (l *XormLoggerAdapter) Warnf(format string, v ...interface{}) {
	logger.Warnf(format, v)
}

// Level implements ContextLogger
func (l *XormLoggerAdapter) Level() log.LogLevel {
	return log.LOG_INFO
}

// SetLevel implements ContextLogger
func (l *XormLoggerAdapter) SetLevel(lv log.LogLevel) {
}

// ShowSQL implements ContextLogger
func (l *XormLoggerAdapter) ShowSQL(show ...bool) {

}

// IsShowSQL implements ContextLogger
func (l *XormLoggerAdapter) IsShowSQL() bool {
	return true
}
