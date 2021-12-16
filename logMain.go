package main

import (
	"github.com/gin-gonic/gin"
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
	"net/http"
)

var serviceLogger *logrus.Logger
var testLogger *logrus.Logger

func main() {
	log.LogPathSet("/Users/zhouzhenyong/tem/tools/logs/tools")
	log.LogApiConfig("/api/core/troy")

	serviceLogger = log.GetLogger("service")
	testLogger = log.GetLogger("test")

	r := gin.Default()
	r.GET("/get", get1)
	r.GET("/test", test)
	r.GET("/service", service)
	log.LogRouters(r)
	log.LogColor(true)
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
		"code":    "success",
		"message": "成功",
		"data":    "12",
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
