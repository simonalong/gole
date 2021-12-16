package main

import (
	"github.com/gin-gonic/gin"
	"github.com/simonalong/tools/log"
	"github.com/simonalong/tools/web"
	"github.com/sirupsen/logrus"
	"net/http"
)

var serviceLogger *logrus.Logger
var testLogger *logrus.Logger

func init() {
	log.LogConfig("/Users/zhouzhenyong/tem/tools/logs/tools", "/api/tools/", true)
}

func main() {

	serviceLogger = log.GetLogger("service")
	testLogger = log.GetLogger("test")

	r := gin.Default()
	//gin.SetMode(gin.ReleaseMode)
	//gin.DefaultWriter = ioutil.Discard
	//gin.DisableConsoleColor()

	r.Use(web.ResponseHandler())
	r.GET("/get", get1)
	r.GET("/test", test)
	r.GET("/service", service)

	// 添加日志api到web
	log.LogRouters(r)

	r.Run(":8082")
}

func get1(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    "12",
	})
}

func test(c *gin.Context) {
	testLogger.Debug("test-debug")
	testLogger.Info("test-debug")
	testLogger.Warn("test-debug")
	testLogger.Error("test-debug")
	//testLogger.Fatalf("test-debug")

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "32",
		"message": "失败",
		"data":    "dfs",
	})
}

func service(c *gin.Context) {
	serviceLogger.Debug("service-debug")
	serviceLogger.Info("service-info")
	serviceLogger.Warn("service-warn")
	serviceLogger.Error("service-error")

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "success",
		"message": "成功",
		"data":    "12",
	})
}