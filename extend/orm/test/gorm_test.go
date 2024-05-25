package test

import (
	"context"
	"fmt"
	"github.com/simonalong/gole/config"
	orm2 "github.com/simonalong/gole/extend/orm"
	"github.com/simonalong/gole/logger"
	"testing"
	"time"
)

func TestGorm1(t *testing.T) {
	config.LoadYamlFile("./application-test1.yaml")
	db, _ := orm2.NewGormDb()

	// 删除表
	db.Exec("drop table isc_demo.gole_demo1")

	//新增表
	db.Exec("CREATE TABLE gole_demo(\n" +
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
	db.Create(&GoleDemo{Name: "zhou", Age: 18, Address: "杭州"})
	db.Create(&GoleDemo{Name: "zhou", Age: 11, Address: "杭州2"})

	// 查询：一行
	var demo GoleDemo
	db.First(&demo).Where("name=?", "zhou")

	dd, _ := db.DB()
	dd.Query("select * from gole_demo")

	// 查询：多行
	fmt.Println(demo)
}

func TestGormOfLoggerChange(t *testing.T) {
	config.LoadYamlFile("./application-test1.yaml")
	//orm2.AddGormHook(&GoleOrmHookDemo{})
	db, _ := orm2.NewGormDb()

	logger.InitLog()

	//// 删除库
	//db.Exec("drop database isc_demo")
	//
	//// 创建库
	//db.Exec("create database isc_demo")

	//新增表
	//db.Exec("CREATE TABLE isc_demo.gole_demo(\n" +
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
	db.Create(&GoleDemo{Name: "zhou", Age: 18, Address: "杭州"})
	db.Create(&GoleDemo{Name: "zhou", Age: 11, Address: "杭州2"})

	// 查询：一行
	var demo GoleDemo
	for i := 0; i < 100; i++ {
		db.First(&demo).Where("name=?", "zhou")
		time.Sleep(time.Second)
		if i == 2 {
			config.SetValue("gole.orm.show-sql", true)
		}

		if i == 4 {
			config.SetValue("gole.orm.show-sql", false)
		}
	}

	// 查询：多行
	fmt.Println(demo)
}

type GoleDemo struct {
	Id      uint64
	Name    string
	Age     int
	Address string
}

func (GoleDemo) TableName() string {
	return "gole_demo"
}

type GoleOrmHookDemo struct {
}

func (*GoleOrmHookDemo) Before(ctx context.Context, driverName string, parameters map[string]any) (context.Context, error) {
	fmt.Println("before")
	fmt.Println(parameters)
	return ctx, nil
}

func (*GoleOrmHookDemo) After(ctx context.Context, driverName string, parameters map[string]any) (context.Context, error) {
	fmt.Println("after")
	fmt.Println(parameters)
	return ctx, nil
}

func (*GoleOrmHookDemo) Err(ctx context.Context, driverName string, err error, parameters map[string]any) error {
	fmt.Println("err")
	fmt.Println(err.Error())
	return nil
}
