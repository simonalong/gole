package log

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"net/http"
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
var unInit = false

var (
	gHost    = "localhost"
	gPort    = "port"
	gApiPath = "/api/tools/"
	gColor   = false
)

// LogConfig fileName 日志文件名
// apiPath 日志管理api路径
// haveColor 日志是否显示颜色
func LogConfig(fileName, apiPath string, haveColor bool) {
	gFilePath = fileName
	if apiPath == "" {
		apiPath = "/api/tools/"
	}

	if !strings.HasSuffix(apiPath, "/") {
		apiPath = apiPath + "/"
	}

	gApiPath = apiPath
	gColor = haveColor
}

func GetLoggerWithConfig(loggerName, filePath, apiPath string, haveColor bool) *logrus.Logger {
	if !unInit {
		LogConfig(filePath, apiPath, haveColor)
		unInit = true
	}
	return GetLogger(loggerName)
}

func GetLogger(loggerName string) *logrus.Logger {
	if logger, exit := loggerMap[loggerName]; exit {
		return logger
	}

	if gFilePath == "" {
		return nil
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
	}, formatters)
	logger.AddHook(lfHook)

	loggerMap[loggerName] = logger
	return logger
}

func LogRouters(r *gin.Engine) {
	appRouter := r.Group(gApiPath)
	{
		// 获取帮助列表
		appRouter.GET("help", getLogToolsHelp)
		// 获取Logger集合curl
		appRouter.GET("logger/list", getLoggerList)
		// 修改host和port
		appRouter.POST("host/change/:host/:port", setHostAndPort)
		// 修改logger的级别
		appRouter.POST("logger/level/:loggerName/:level", setLoggerLevel)
		// 修改所有logger的级别
		appRouter.POST("logger/root/level/:level", setLoggerRootLevel)
	}
}

func getLogToolsHelp(c *gin.Context) {
	helpStr := map[string]string{}
	helpStr["查询：帮助-----------"] = fmt.Sprintf("curl %v%vhelp", getHostAndPort(), gApiPath)
	helpStr["查询：Logger集合-----"] = fmt.Sprintf("curl %v%vlogger/list", getHostAndPort(), gApiPath)
	helpStr["修改：host和port-----"] = fmt.Sprintf("curl -X POST %v%vhost/change/{host}/{port}", getHostAndPort(), gApiPath)
	helpStr["修改：logger的级别----"] = fmt.Sprintf("curl -X POST %v%vlogger/level/{loggerName}/{level}", getHostAndPort(), gApiPath)
	helpStr["修改：所有logger的级别"] = fmt.Sprintf("curl -X POST %v%vlogger/root/level/{level}", getHostAndPort(), gApiPath)
	Success(c, helpStr)
}

func getLoggerList(c *gin.Context) {
	var keys []string
	for key, _ := range loggerMap {
		keys = append(keys, key)
	}
	Success(c, keys)
}

func setLoggerLevel(c *gin.Context) {
	loggerName := c.Param("loggerName")
	level := c.Param("level")
	if loggerValue, exist := loggerMap[loggerName]; exist {
		levelValue, err := logrus.ParseLevel(strings.ToLower(level))
		if err != nil {
			return
		}
		loggerValue.SetLevel(levelValue)
	}
	Success(c, 1)
}

func setHostAndPort(c *gin.Context) {
	host := c.Param("host")
	port := c.Param("port")
	gHost = host
	gPort = port
}

func setLoggerRootLevel(c *gin.Context) {
	level := c.Param("level")
	for _, logger := range loggerMap {
		levelValue, err := logrus.ParseLevel(strings.ToLower(level))
		if err != nil {
			return
		}
		logger.SetLevel(levelValue)
	}
	Success(c, len(loggerMap))
}

func getHostAndPort() string {
	return "http://" + gHost + ":" + gPort
}

func rotateLog(path, level string) *rotatelogs.RotateLogs {
	if pRotateValue, exist := rotateMap[path+"-"+level]; exist {
		return pRotateValue
	}

	if rotateMap == nil {
		rotateMap = map[string]*rotatelogs.RotateLogs{}
	}

	if path == "" {
		path = "./logs/app"
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
	var levelColor = gray
	switch level {
	case logrus.DebugLevel:
		if gColor {
			levelColor = blue
		}
	case logrus.InfoLevel:
		if gColor {
			levelColor = blue
		}
	case logrus.WarnLevel:
		if gColor {
			levelColor = blue
		}
	case logrus.ErrorLevel:
		if gColor {
			levelColor = red
		}
	}

	if gColor {
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s %s\n", levelColor, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, fieldsStr)
	} else {
		newLog = fmt.Sprintf("%s\t %s %s %s %s\n", strings.ToUpper(entry.Level.String()), timestamp, funPath, entry.Message, fieldsStr)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func Success(ctx *gin.Context, object interface{}) {
	ctx.JSON(http.StatusOK, object)
}
