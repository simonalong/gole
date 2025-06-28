package test

import (
	orm2 "github.com/simonalong/gole/extend/orm"
	"github.com/simonalong/gole/extend/orm/test/do"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/config"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/excel"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/logger"
	"gitlab.seatakcloud.com/cbb/base/cbb-base/util"
	"testing"
)

// 查询一行
func TestSelectOne(t *testing.T) {
	config.LoadYamlFile("./application.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}
	user := do.User{}
	// 获取的就是id最小的一行
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	gormClient.First(&user)
	logger.Info(user)

	user2 := do.User{}
	// 这个没有orderBy，但是获取到的是第一条
	// SELECT * FROM `users` LIMIT 1
	gormClient.Take(&user2)
	logger.Info(user2)

	user3 := do.User{}
	// 这个没有orderBy，但是获取到的是第一条
	// SELECT * FROM `users` ORDER BY `users`.`id`,`users`.`id` DESC LIMIT 1
	gormClient.Last(&user3)
	logger.Info(user3)
}

func TestSelectCount(t *testing.T) {
	config.LoadYamlFile("./application-rds.yaml")
	gormClient, err := orm2.NewGormClient()
	if err != nil {
		logger.Fatalf("连接数据库异常：%v", err)
	}

	datas, err := excel.ReadDatas("./数据库数据统计.xlsx")
	if err != nil {
		logger.Fatalf("读取数据异常：%v", err)
	}

	for rowId, row := range datas {
		if rowId <= 197 {
			continue
		}
		var count int64
		gormClient.Table(row[0]).Count(&count)
		logger.Info("表名：", row[0], "；行数：", count)
		row[1] = util.ToString(count)
	}
	excel.AddDatas("./数据库数据统计_new.xlsx", datas)
}
