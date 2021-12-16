package tools

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lunny/log"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"path/filepath"
	"strings"
	"time"
)

const (
	white  = 29
	black  = 30
	red    = 31
	green  = 32
	yellow = 33
	purple = 35
	blue   = 36
	gray   = 37
)

var loggerMap map[string]*logrus.Logger
var rotateMap map[string]*rotatelogs.RotateLogs
var gFilePath string

func LogPathSet(fileName string) {
	gFilePath = fileName
}

func GetLogger(loggerName string) *logrus.Logger {
	if logger, exit := loggerMap[loggerName]; exit {
		return logger
	}

	if gFilePath == "" {
		log.Errorf("please set file path")
	}

	if loggerMap == nil {
		loggerMap = map[string]*logrus.Logger{}
	}
	logger := logrus.New()

	logger.SetReportCaller(true)
	formatters := &StandardFormatter{}
	logger.Formatter = formatters

	lfHook := lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: rotateLog(gFilePath, "debug"),
		logrus.InfoLevel:  rotateLog(gFilePath, "info"),
		logrus.WarnLevel:  rotateLog(gFilePath, "warn"),
		logrus.ErrorLevel: rotateLog(gFilePath, "error"),
		logrus.FatalLevel: rotateLog(gFilePath, "fatal"),
		logrus.PanicLevel: rotateLog(gFilePath, "panic"),
	}, formatters)
	logger.AddHook(lfHook)

	loggerMap[loggerName] = logger
	return logger
}

func rotateLog(path, level string) *rotatelogs.RotateLogs {
	if pRotateValue, exist := rotateMap[path+"-"+level]; exist {
		return pRotateValue
	}

	if rotateMap == nil {
		rotateMap = map[string]*rotatelogs.RotateLogs{}
	}

	data, _ := rotatelogs.New(path+"-"+level+".log.%Y%m%d", rotatelogs.WithLinkName(path+"-"+level+".log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
	rotateMap[path+"-"+level] = data
	return data
}

type StandardFormatter struct{}

func (m *StandardFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	var fields []string
	for k, v := range entry.Data {
		fields = append(fields, fmt.Sprintf("%v=%v", k, v))
	}

	level := entry.Level
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	var funPath string
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		funPath = fmt.Sprintf("%s %s:%d", entry.Caller.Function, fName, entry.Caller.Line)
	} else {
		funPath = fmt.Sprintf("%s", entry.Message)
	}

	var fieldsStr string
	if len(fields) != 0 {
		fieldsStr = fmt.Sprintf("[\x1b[%dm%s\x1b[0m]", blue, strings.Join(fields, " "))
	}
	var newLog string
	switch level {
	case logrus.DebugLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s %s\n", white, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, fieldsStr)
	case logrus.InfoLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s %s\n", green, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, fieldsStr)
	case logrus.WarnLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s %s\n", yellow, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, fieldsStr)
	case logrus.ErrorLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s %s\n", red, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, fieldsStr)
	case logrus.FatalLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s %s\n", purple, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, fieldsStr)
	case logrus.PanicLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s %s", blue, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, fieldsStr)
	}
	b.WriteString(newLog)

	return b.Bytes(), nil
}
