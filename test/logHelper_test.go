package test

import (
	"github.com/gin-gonic/gin"
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
	"net/http"
	"testing"
)

func TestLoggerGet(t *testing.T) {
	serviceLog := log.GetLogger("service")
	serviceLog.Info("haode")
	serviceLog.SetLevel(logrus.DebugLevel)
	serviceLog.WithField("nihao", 12).Debug("haode")
	serviceLog.WithField("nihao", 12).Info("haode")
	serviceLog.WithField("nihao", 12).Warn("haode")
	serviceLog.WithField("nihao", 12).Error("haode")
	serviceLog.WithField("nihao", 12).Fatal("haode")
}

func TestLogRouter(t *testing.T) {
	var engine = gin.Default()
	log.LogRouters(engine)
	engine.GET("/handle", haneld)
	engine.Run(":8090")
}

func haneld(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
