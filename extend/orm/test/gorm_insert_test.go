package test

import (
	"context"
	"fmt"
	"github.com/simonalong/gole/config"
	orm2 "github.com/simonalong/gole/extend/orm"
	"github.com/simonalong/gole/extend/orm/test/do"
	"github.com/simonalong/gole/logger"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"testing"
	"time"
)

// 基本功能
// 1.增加
// 2.删除
// 3.修改
// 4.查询
// -- 1.单数据查询
// -- 2.列表数据查询
// -- 3.分页数据查询
// -- 4.个数的查询
// -- 5.直接执行sql

func TestInsert(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	gormClient.AutoMigrate(do.User{})

	dir := time.Now()
	user := do.User{Name: "Jinzhu", Age: 18, Birthday: &dir}

	rows := gormClient.Create(&user).RowsAffected // 通过数据的指针来创建
	logger.Infof("rows: %v；err: %v", rows, err)
}

// 指定某个属性插入：使用Select方法
func TestInsert2(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	gormClient.AutoMigrate(do.User{})

	dir := time.Now()
	user := do.User{Name: "Jinzhu", Age: 18, Birthday: &dir}

	// INSERT INTO `users` (`age`) VALUES (18)
	rows := gormClient.Select("age").Create(&user).RowsAffected // 通过数据的指针来创建
	logger.Infof("rows: %v；err: %v, user.id=%v", rows, err, user.ID)
}

// 指定某几个字段不插入：使用Omit方法
func TestInsert3(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	gormClient.AutoMigrate(do.User{})

	dir := time.Now()
	user := do.User{Name: "Jinzhu", Age: 18, Birthday: &dir}

	// INSERT INTO `users` (`name`,`email`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('Jinzhu',NULL,'2025-01-25 13:35:27.375',NULL,NULL,'2025-01-25 13:35:27.376','2025-01-25 13:35:27.376')
	rows := gormClient.Omit("age").Create(&user).RowsAffected // 通过数据的指针来创建
	logger.Infof("rows: %v；err: %v, user.id=%v", rows, err, user.ID)
}

func TestBatchInsert(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	dir := time.Now()
	users := []do.User{
		{Name: "Jinzhu1", Age: 11, Birthday: &dir},
		{Name: "Jinzhu2", Age: 12, Birthday: &dir},
		{Name: "Jinzhu3", Age: 13, Birthday: &dir},
	}

	// INSERT INTO `users` (`name`,`email`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('Jinzhu',NULL,'2025-01-25 13:35:27.375',NULL,NULL,'2025-01-25 13:35:27.376','2025-01-25 13:35:27.376')
	rows := gormClient.Create(&users).RowsAffected // 通过数据的指针来创建
	for _, user := range users {
		logger.Infof("rows: %v；err: %v, user.id=%v", rows, err, user.ID)
	}
}

// 使用map结构来插入数据
func TestInsert4_Map(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	dir := time.Now()

	rows := gormClient.Model(&do.User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18, "birthday": dir,
	}).RowsAffected
	logger.Infof("rows: %v；err: %v", rows, err)
}

func TestBatchInsert_Map(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	dir := time.Now()

	rows := gormClient.Model(&do.User{}).Create([]map[string]interface{}{
		{"Name": "jinzhu1", "Age": 11, "birthday": dir},
		{"Name": "jinzhu2", "Age": 12, "birthday": dir},
		{"Name": "jinzhu3", "Age": 13, "birthday": dir},
	}).RowsAffected
	logger.Infof("rows: %v；err: %v", rows, err)
}

func TestBatchInsert5(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	gormClient.AutoMigrate(UserData{})

	// Create from map
	gormClient.Model(UserData{}).Create(map[string]interface{}{
		"Name":     "jinzhu",
		"Location": clause.Expr{SQL: "ST_PointFromText(?)", Vars: []interface{}{"POINT(100 100)"}},
	})

	// INSERT INTO `users` (`name`,`location`) VALUES ("jinzhu",ST_PointFromText("POINT(100 100)"));
	gormClient.Create(&UserData{
		Name:     "jinzhu",
		Location: Location{X: 100, Y: 100},
	})
}

// Create from customized data type
type Location struct {
	X, Y int
}

// Scan implements the sql.Scanner interface
func (loc *Location) Scan(v interface{}) error {
	// Scan a value into struct from database driver
	return nil
}

func (loc Location) GormDataType() string {
	return "geometry"
}

func (loc Location) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%d %d)", loc.X, loc.Y)},
	}
}

type UserData struct {
	Name     string
	Location Location
}

type UserDemo6 struct {
	ID   int64
	Name string `gorm:"default:galeone"`
	Age  int64  `gorm:"default:18"`
}

func TestInsert6(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	gormClient.AutoMigrate(UserDemo6{})

	rows := gormClient.Create(&UserDemo6{
		Name: "zhou", Age: 12,
	}).RowsAffected

	rows = gormClient.Create(&UserDemo6{}).RowsAffected
	// 这里的age是设置不了为0的，因为有默认值，这个默认值触发的条件就是这个age为0；如果想要设置这个age为0，则请看例子7下面这个Insert7
	rows = gormClient.Create(&UserDemo6{Age: 0}).RowsAffected
	logger.Infof("rows: %v；err: %v", rows, err)
}

type UserDemo7 struct {
	ID   int64
	Name string
	Age  *int64 `gorm:"default:18"`
}

func TestInsert7(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	gormClient.AutoMigrate(UserDemo7{})

	// 这样插入的这个age就是0，而不是默认值18
	var age int64 = 0
	rows := gormClient.Create(&UserDemo7{Age: &age}).RowsAffected
	logger.Infof("rows: %v；err: %v", rows, err)
}

// 在数据插入时候出现冲突的情况下，应该的处理方式
func TestUpset(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	// 这样插入的这个age就是0，而不是默认值18
	rows := gormClient.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"age": 12}),
	}).Create(&do.User{ID: 13, Age: 2}).RowsAffected
	logger.Infof("rows: %v；err: %v", rows, err)
}
