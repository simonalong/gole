package test

import (
	"context"
	"fmt"
	"github.com/simonalong/gole/config"
	orm2 "github.com/simonalong/gole/extend/orm"
	"testing"
	"xorm.io/xorm/contexts"
)

func TestXorm1(t *testing.T) {
	config.LoadYamlFile("./application-test1.yaml")
	orm2.AddXormHook(&BaseXormHook{})
	xormClient, _ := orm2.NewXormClient()

	// 删除表
	xormClient.Exec("drop table test.base_demo")

	//新增表
	xormClient.Exec("CREATE TABLE base_demo(\n" +
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

	xormClient.Table("base_demo").Insert(&BaseDemo{Name: "zhou", Age: 18, Address: "杭州"})
	// 新增
	xormClient.Table("base_demo").Insert(&BaseDemo{Name: "zhou", Age: 18, Address: "杭州"})

	var demo BaseDemo
	xormClient.Table("base_demo").Where("name=?", "zhou").Get(&demo)

	dd := xormClient.DB()
	dd.Query("select * from base_demo")

	// 查询：多行
	fmt.Println(demo)
}

type BaseXormHook struct {
}

func (*BaseXormHook) BeforeProcess(c *contexts.ContextHook, driverName string) (context.Context, error) {
	fmt.Println("before-xorm")
	return c.Ctx, nil
}

func (*BaseXormHook) AfterProcess(c *contexts.ContextHook, driverName string) error {
	fmt.Println("after-xorm")
	return nil
}
