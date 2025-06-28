package logger

import (
	"bytes"
	"fmt"
	"github.com/gookit/color"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/rifflock/lfshook"
	"github.com/simonalong/gole/config"
	"github.com/simonalong/gole/util"
	"github.com/sirupsen/logrus"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

const (
	maximumCallerDepth    int = 25
	knownBaseLoggerFrames int = 5
)

var callerInitOnce sync.Once
var minimumCallerDepth = 0
var baseLoggerPackage string

var gColor = true

// var loggerMap map[string]*logrus.Logger
var loggerMap cmap.ConcurrentMap

// var rotateMap map[string]*rotatelogs.RotateLogs
var rotateMap cmap.ConcurrentMap
var rootLogger *logrus.Logger

func init() {
	loggerMap = cmap.New()
	rotateMap = cmap.New()
	rootLogger = Group("root")

	_gColor := config.GetValueBoolDefault("base.logger.color.enable", false)
	gColor = _gColor

	InitLog()
}

func GetRootLogger() *logrus.Logger {
	return rootLogger
}

func SetGroupLevel(groupName, level string) {
	if level == "" {
		return
	}
	level = strings.ToLower(level)
	switch level {
	case "debug":
		Group(groupName).SetLevel(logrus.DebugLevel)
	case "info":
		Group(groupName).SetLevel(logrus.InfoLevel)
	case "warn":
		Group(groupName).SetLevel(logrus.WarnLevel)
	case "error":
		Group(groupName).SetLevel(logrus.ErrorLevel)
	case "fatal":
		Group(groupName).SetLevel(logrus.FatalLevel)
	}
}

func Group(groupNames ...string) *logrus.Logger {
	var resultLogger *logrus.Logger
	var groupNamesOfUnContain []string
	for _, groupName := range groupNames {
		if logger, exit := loggerMap.Get(groupName); exit {
			resultLogger = logger.(*logrus.Logger)
		} else {
			groupNamesOfUnContain = append(groupNamesOfUnContain, groupName)
		}
	}

	if resultLogger != nil {
		return resultLogger
	} else {
		resultLogger = logrus.New()
		resultLogger.SetReportCaller(true)
		formatters := &StandardFormatter{}
		resultLogger.Formatter = formatters

		loggerDir := config.GetValueStringDefault("base.logger.home", "."+string(os.PathSeparator)+"logs"+string(os.PathSeparator))
		resultLogger.AddHook(lfshook.NewHook(lfshook.WriterMap{
			logrus.DebugLevel: rotateLogWithCache(loggerDir, "debug"),
			logrus.InfoLevel:  rotateLogWithCache(loggerDir, "info"),
			logrus.WarnLevel:  rotateLogWithCache(loggerDir, "warn"),
			logrus.ErrorLevel: rotateLogWithCache(loggerDir, "error"),
			logrus.PanicLevel: rotateLogWithCache(loggerDir, "panic"),
			logrus.FatalLevel: rotateLogWithCache(loggerDir, "fatal"),
		}, formatters))
	}

	// 值最大的级别，对应的level最小，比如Debug对应的数值比Info要大
	maxValueLevel := logrus.PanicLevel
	for _, groupName := range groupNamesOfUnContain {
		var finalGroupLevel string
		rootLevel := config.GetValueStringDefault("base.logger.level", "info")
		groupLevel := config.GetValueString("base.logger.group." + groupName + ".level")
		if groupLevel != "" {
			finalGroupLevel = groupLevel
		} else {
			finalGroupLevel = rootLevel
		}

		lgLevel, err := logrus.ParseLevel(finalGroupLevel)
		if err != nil {
			lgLevel = logrus.InfoLevel
		}

		if lgLevel > maxValueLevel {
			maxValueLevel = lgLevel
		}
	}

	resultLogger.SetLevel(maxValueLevel)

	for _, groupName := range groupNamesOfUnContain {
		loggerMap.Set(groupName, resultLogger)
	}
	return resultLogger
}

func doGroup(groupName string) *logrus.Logger {
	if groupName == "" {
		return rootLogger
	}
	if logger, exit := loggerMap.Get(groupName); exit {
		return logger.(*logrus.Logger)
	}

	if loggerMap == nil {
		loggerMap = cmap.New()
	}
	logger := logrus.New()
	logger.SetReportCaller(true)
	formatters := &StandardFormatter{}
	logger.Formatter = formatters

	loggerDir := config.GetValueStringDefault("base.logger.home", "./logs/")
	logger.AddHook(lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: rotateLogWithCache(loggerDir, "debug"),
		logrus.InfoLevel:  rotateLogWithCache(loggerDir, "info"),
		logrus.WarnLevel:  rotateLogWithCache(loggerDir, "warn"),
		logrus.ErrorLevel: rotateLogWithCache(loggerDir, "error"),
		logrus.PanicLevel: rotateLogWithCache(loggerDir, "panic"),
		logrus.FatalLevel: rotateLogWithCache(loggerDir, "fatal"),
	}, formatters))

	var finalGroupLevel string
	rootLevel := config.GetValueStringDefault("base.logger.level", "info")
	groupLevel := config.GetValueString("base.logger.group." + groupName + ".level")
	if groupLevel != "" {
		finalGroupLevel = groupLevel
	} else {
		finalGroupLevel = rootLevel
	}

	lgLevel, err := logrus.ParseLevel(finalGroupLevel)
	if err != nil {
		lgLevel = logrus.InfoLevel
	}
	logger.SetLevel(lgLevel)

	loggerMap.Set(groupName, logger)
	return logger
}

func InitLog() {
	rootLogger = Group("root")
	loggerDir := config.GetValueStringDefault("base.logger.home", "./logs/")
	rootLogger.AddHook(lfshook.NewHook(lfshook.WriterMap{
		logrus.DebugLevel: rotateLog(loggerDir, "debug"),
		logrus.InfoLevel:  rotateLog(loggerDir, "info"),
		logrus.WarnLevel:  rotateLog(loggerDir, "warn"),
		logrus.ErrorLevel: rotateLog(loggerDir, "error"),
		logrus.PanicLevel: rotateLog(loggerDir, "panic"),
		logrus.FatalLevel: rotateLog(loggerDir, "fatal"),
	}, &StandardFormatter{}))
	lgLevel, err := logrus.ParseLevel(config.GetValueStringDefault("base.logger.level", "info"))
	if err != nil {
		lgLevel = logrus.InfoLevel
	}
	rootLogger.SetLevel(lgLevel)

	_gColor := config.GetValueBoolDefault("base.logger.color.enable", false)
	gColor = _gColor
}

func GetLoggerGroupList(name string) []string {
	if name == "_all_" {
		return loggerMap.Keys()
	}
	var groupNames []string
	for _, key := range loggerMap.Keys() {
		if strings.Contains(key, name) {
			groupNames = append(groupNames, key)
		}
	}
	return groupNames
}

func SetGlobalLevel(strLevel string) {
	level, err := logrus.ParseLevel(strLevel)
	if err == nil {
		rootLogger.SetLevel(level)
	}
}

// Trace 这个日志级别不建议使用，因为目前来看没有日志输出
// Deprecated
func Trace(v ...any) {
	rootLogger.Trace(v...)
}

func Debug(v ...any) {
	rootLogger.Debug(v...)
}

func Info(v ...any) {
	rootLogger.Info(v...)
}

func Warn(v ...any) {
	rootLogger.Warn(v...)
}

func Error(v ...any) {
	rootLogger.Error(v...)
}

func Fatal(v ...any) {
	rootLogger.Fatal(v...)
}

// Panic 这个日志级别也不建议使用
// Deprecated
func Panic(v ...any) {
	rootLogger.Panic(v...)
}

// Tracef 这个日志级别也不建议使用
// Deprecated
func Tracef(format string, v ...any) {
	rootLogger.Tracef(format, v...)
}

func Debugf(format string, v ...any) {
	rootLogger.Debugf(format, v...)
}

func Infof(format string, v ...any) {
	rootLogger.Infof(format, v...)
}

func Warnf(format string, v ...any) {
	rootLogger.Warnf(format, v...)
}

func Errorf(format string, v ...any) {
	rootLogger.Errorf(format, v...)
}

// Panicf 不建议使用
// Deprecated
func Panicf(format string, v ...any) {
	rootLogger.Panicf(format, v...)
}

func Fatalf(format string, v ...any) {
	rootLogger.Fatalf(format, v...)
}

func Record(level, format string, v ...any) {
	switch strings.ToLower(level) {
	case "debug":
		Debugf(format, v)
	case "info":
		Infof(format, v)
	case "warn":
		Warnf(format, v)
	case "error":
		Errorf(format, v)
	case "panic":
		Panicf(format, v)
	case "fatal":
		Fatalf(format, v)
	default:
		Debugf(format, v)
	}
}

func rotateLog(path, level string) *rotatelogs.RotateLogs {
	if rotateMap == nil {
		rotateMap = cmap.New()
	}

	if path == "" {
		path = "." + string(os.PathSeparator) + "logs" + string(os.PathSeparator)
	}

	maxSizeStr := config.GetValueStringDefault("base.logger.rotate.max-size", "300MB")
	maxHistoryStr := config.GetValueStringDefault("base.logger.rotate.max-history", "60d")
	rotateTimeStr := config.GetValueStringDefault("base.logger.rotate.time", "1d")

	rotateOptions := []rotatelogs.Option{rotatelogs.WithLinkName(path + level + ".log")}
	if maxSizeStr != "" {
		rotateOptions = append(rotateOptions, rotatelogs.WithRotationSize(util.ParseByteSize(maxSizeStr)))
	}

	_maxHistory, err := time.ParseDuration(maxHistoryStr)
	if err == nil {
		rotateOptions = append(rotateOptions, rotatelogs.WithMaxAge(_maxHistory))
	}

	_rotateTime, err := time.ParseDuration(rotateTimeStr)
	if err == nil {
		rotateOptions = append(rotateOptions, rotatelogs.WithRotationTime(_rotateTime))
	}

	data, _ := rotatelogs.New(path+level+".%Y%m%d.log", rotateOptions...)
	rotateMap.Set(path+"-"+level, data)
	return data
}

func rotateLogWithCache(path, level string) *rotatelogs.RotateLogs {
	if pRotateValue, exist := rotateMap.Get(path + "-" + level); exist {
		return pRotateValue.(*rotatelogs.RotateLogs)
	}

	return rotateLog(path, level)
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

	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	var funPath string
	if entry.HasCaller() {
		frame := getCallerFrame()
		funPath = fmt.Sprintf("%s:%d#%s", shortLogPath(frame.File), frame.Line, functionName(frame))
	} else {
		funPath = fmt.Sprintf("%s", entry.Message)
	}

	var fieldsStr string
	if len(fields) != 0 {
		fieldsStr = strings.Join(fields, " ")
	}
	var newLog string

	// todo 租户id，字段先预留
	var tenantId string
	if gColor {
		levelStr := levelToStr(entry.Level)
		newLog = fmt.Sprintf("%s[%s][%s][tid=%s][%s][%v] %s %s\n",
			colorTimestampStr(levelStr, timestamp),
			color.FgDarkGray.Render(config.GetValueStringDefault("base.application.name", "base")),
			colorLevelStr(levelStr),
			color.FgLightCyan.Render(""),
			tenantId,
			color.FgDarkGray.Render(funPath),
			colorMsgStr(levelStr, entry.Message),
			color.FgCyan.Render(fieldsStr))
	} else {
		newLog = fmt.Sprintf("%s[%s][%s][tid=%s][%s][%v] %s %s\n",
			timestamp,
			config.GetValueStringDefault("base.application.name", "base"),
			levelToStr(entry.Level),
			"",
			tenantId,
			funPath,
			entry.Message,
			fieldsStr)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, string(os.PathSeparator))
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}
	return f
}

func levelToStr(level logrus.Level) string {
	levelStr := strings.ToUpper(level.String())
	if levelStr == "WARNING" {
		levelStr = "WARN"
	}
	return levelStr
}

func getCallerFrame() *runtime.Frame {
	pcs := make([]uintptr, maximumCallerDepth)
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "logger.getCallerFrame") {
				baseLoggerPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownBaseLoggerFrames
	})

	pcs = make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)
		if pkg != baseLoggerPackage && pkg != "github.com/sirupsen/logrus" {
			return &f
		}
	}
	return nil
}

func functionName(frame *runtime.Frame) string {
	pathMeta := strings.Split(frame.Function, ".")
	if len(pathMeta) > 1 {
		return pathMeta[len(pathMeta)-1]
	}
	return frame.Function
}

func shortLogPath(logPath string) string {
	loggerPath := config.GetValueStringDefault("base.logger.path.type", "short")
	if loggerPath == "short" {
		pathMeta := strings.Split(logPath, string(os.PathSeparator))
		if len(pathMeta) > 3 {
			return pathMeta[len(pathMeta)-3] + string(os.PathSeparator) + pathMeta[len(pathMeta)-2] + string(os.PathSeparator) + pathMeta[len(pathMeta)-1]
		}
		return logPath
	} else if loggerPath == "full" {
		pathMeta := strings.Split(logPath, "@2/project")
		if len(pathMeta) > 1 {
			pathMeta[0] = "../.."
			return strings.Join(pathMeta, "")
		}
		return logPath
	} else {
		return logPath
	}
}
