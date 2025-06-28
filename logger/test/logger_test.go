package test

import (
	"github.com/simonalong/gole/config"
	"github.com/simonalong/gole/logger"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInfo1(t *testing.T) {
	config.LoadYamlFile("./application-debug.yaml")
	////logger.InitLog()

	logger.Info("hello ", "info")
	logger.Infof("hello %v", "info")
}

func TestInfo(t *testing.T) {
	config.LoadYamlFile("./application-debug.yaml")
	//logger.InitLog()

	// info
	logger.Infof("hello %v", "info")
	logger.Group("group1").Info("hello", " ", "info")
	logger.Group("group1").Infof("hello %v", "info")
}

func TestLevel(t *testing.T) {
	config.LoadYamlFile("./application-debug.yaml")
	//logger.InitLog()

	// debug
	logger.Debugf("hello %v", "debug")

	// info
	logger.Infof("hello %v", "info")

	// warn
	logger.Warnf("hello %v", "warn")

	// error
	logger.Errorf("hello %v", "error")
}

func TestLevelChange(t *testing.T) {
	config.LoadYamlFile("./application-debug.yaml")
	//logger.InitLog()

	// info
	logger.Infof("hello %v", "info1")
	logger.Infof("hello %v", "info2")

	// 设置后下面的不再显示
	logger.SetGlobalLevel("warn")
	logger.Infof("hello %v", "info3")
}

// 日志分组的级别变更
func TestGroupLevelChange1(t *testing.T) {
	config.LoadYamlFile("./application-debug.yaml")
	//logger.InitLog()

	// info
	logger.Infof("hello %v", "info1")
	logger.Group("group1").Infof("hello %v", "group1 info1")
	logger.Group("group2").Infof("hello %v", "group2 info2")

	// 设置后下面的不再显示
	logger.SetGlobalLevel("warn")
	// 不再打印
	logger.Infof("hello %v", "info3")
	// 继续打印
	logger.Group("group1").Infof("hello %v", "group1 info1")
	logger.Group("group2").Infof("hello %v", "group2 info2")

	warnLevel, _ := logrus.ParseLevel("warn")
	logger.Group("group1").SetLevel(warnLevel)
	// 不再打印
	logger.Group("group1").Infof("hello %v", "group1 change info1")
	// 继续打印
	logger.Group("group2").Infof("hello %v", "group2 change info2")
}

// 日志分组的级别变更
func TestGroupLevelChange2(t *testing.T) {
	config.LoadYamlFile("./application-group.yaml")
	//logger.InitLog()

	logger.SetGlobalLevel("error")
	logger.SetGroupLevel("g1", "debug")
	logger.SetGroupLevel("g2", "error")

	// info
	logger.Debugf("hello %v", "debug")
	logger.Infof("hello %v", "info")
	logger.Warnf("hello %v", "warn")

	// 只有g1的打印
	logger.Group("g1").Debugf("hello %v", "g1 debug")
	logger.Group("g1").Infof("hello %v", "g1 info")
	logger.Group("g1").Warnf("hello %v", "g1 warn")

	logger.Group("g2").Debugf("hello %v", "g2 debug")
	logger.Group("g2").Infof("hello %v", "g2 info")
	logger.Group("g2").Warnf("hello %v", "g2 warn")
}

func TestLoggerPathShort(t *testing.T) {
	config.LoadYamlFile("./application-short.yaml")
	//logger.InitLog()

	logger.Info("test")
}

func TestLoggerPathFull(t *testing.T) {
	config.LoadYamlFile("./application-full.yaml")
	//logger.InitLog()

	logger.Info("test")
}

func TestLoggerRotate(t *testing.T) {
	config.LoadYamlFile("./application-rotate.yaml")
	//logger.InitLog()

	//for i := 0; i < 100; i++ {
	//	logger.Info("test")
	//	time.Sleep(1 * time.Second)
	//}
}

func TestLoggerGroup2(t *testing.T) {
	config.LoadYamlFile("./application-group2.yaml")
	//logger.InitLog()

	logger.Group("g1", "g2").Debug("test")
}
