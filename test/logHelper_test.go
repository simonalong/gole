package test

import (
	"github.com/simonalong/tools"
	"testing"
)

func TestLoggerGet(t *testing.T) {
	tools.LogInit("/Users/zhouzhenyong/tem/tools/app")
	serviceLog := tools.LoggerGet("service")
	serviceLog.Info("haode")
}
