package test

import (
	"context"
	"fmt"
	orm2 "github.com/simonalong/gole/extend/orm"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/config"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/logger"
	"testing"
	"time"
)

func TestGorm1(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	db, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("数据库连接创建失败：%v", err)
		return
	}

	// 删除表
	db.Exec("drop table test.base_demo")

	//新增表
	db.Exec("CREATE TABLE base_demo(\n" +
		"  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',\n" +
		"  `name` char(20) NOT NULL COMMENT '名字',\n" +
		"  `age` INT NOT NULL COMMENT '年龄',\n" +
		"  `address` char(20) NOT NULL COMMENT '名字',\n" +
		"  \n" +
		"  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
		"  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
		"\n" +
		"  PRIMARY KEY (`id`)\n" +
		") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试表'")

	// 新增
	db.Create(&BaseDemo{Name: "zhou", Age: 18, Address: "杭州"})
	db.Create(&BaseDemo{Name: "zhou", Age: 11, Address: "杭州2"})

	// 查询：一行
	var demo BaseDemo
	db.First(&demo).Where("name=?", "zhou")

	dd, _ := db.DB()
	dd.Query("select * from base_demo")

	// 查询：多行
	fmt.Println(demo)
}

func TestGormOfLoggerChange(t *testing.T) {
	config.LoadYamlFile("./application-test1.yaml")
	//orm2.AddGormHook(&MeterGormHook{})
	db, _ := orm2.NewGormClient()

	logger.InitLog()

	//// 删除库
	//db.Exec("drop database test")
	//
	//// 创建库
	//db.Exec("create database test")

	//新增表
	//db.Exec("CREATE TABLE test.base_demo(\n" +
	//	"  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键',\n" +
	//	"  `name` char(20) NOT NULL COMMENT '名字',\n" +
	//	"  `age` INT NOT NULL COMMENT '年龄',\n" +
	//	"  `address` char(20) NOT NULL COMMENT '名字',\n" +
	//	"  \n" +
	//	"  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',\n" +
	//	"  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',\n" +
	//	"\n" +
	//	"  PRIMARY KEY (`id`)\n" +
	//	") ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='测试表'")

	// 新增
	db.Create(&BaseDemo{Name: "zhou", Age: 18, Address: "杭州"})
	db.Create(&BaseDemo{Name: "zhou", Age: 11, Address: "杭州2"})

	// 查询：一行
	var demo BaseDemo
	for i := 0; i < 100; i++ {
		db.First(&demo).Where("name=?", "zhou")
		time.Sleep(time.Second)
		if i == 2 {
			config.SetValue("base.orm.show-sql", true)
		}

		if i == 4 {
			config.SetValue("base.orm.show-sql", false)
		}
	}

	// 查询：多行
	fmt.Println(demo)
}

type BaseDemo struct {
	Id      uint64
	Name    string
	Age     int
	Address string
}

func (BaseDemo) TableName() string {
	return "base_demo"
}

type BaseOrmHookDemo struct {
}

func (*BaseOrmHookDemo) Before(ctx context.Context, driverName string, parameters map[string]any) (context.Context, error) {
	fmt.Println("before")
	fmt.Println(parameters)
	return ctx, nil
}

func (*BaseOrmHookDemo) After(ctx context.Context, driverName string, parameters map[string]any) (context.Context, error) {
	fmt.Println("after")
	fmt.Println(parameters)
	return ctx, nil
}

func (*BaseOrmHookDemo) Err(ctx context.Context, driverName string, err error, parameters map[string]any) error {
	fmt.Println("err")
	fmt.Println(err.Error())
	return nil
}

func TestGormHook(t *testing.T) {
	config.LoadYamlFile("./application-test1.yaml")
	orm2.AddGormHook(&BaseOrmHookDemo{})
	db, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("数据库连接创建失败：%v", err)
		return
	}

	var demo BaseDemo
	db.First(&demo).Where("name=?", "zhou")
	fmt.Println(demo)
}
