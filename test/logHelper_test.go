package test

import (
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
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
