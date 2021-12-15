package tools

import (
	"bytes"
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/lunny/log"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io"
	"runtime"
	"strings"
	"sync"
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

const (
	maximumCallerDepth int = 25
	knownLogrusFrames  int = 4
)

var loggerMap map[string]*logrus.Logger
var fileMap map[string]map[logrus.Level]io.Writer
var gFilePath string

func InitLogPath(fileName string) {
	rotateMap(fileName)
	gFilePath = fileName
}

func GetLogger(loggerName string) *logrus.Logger {
	if logger, exit := loggerMap[loggerName]; exit {
		return logger
	}

	if gFilePath == "" {
		log.Errorf("please set file path")
	}
	logger := logrus.New()
	logger.Formatter = &StandardFormatter{}
	logger.AddHook(lfshook.NewHook(rotateMap(gFilePath), &StandardFormatter{}))

	loggerMap[loggerName] = logger
	return logger
}

func rotateMap(fileName string) map[logrus.Level]io.Writer {
	if rotate, exist := fileMap[fileName]; exist {
		return rotate
	}
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: rotateLog(fileName, "debug"),
		logrus.InfoLevel:  rotateLog(fileName, "info"),
		logrus.WarnLevel:  rotateLog(fileName, "warn"),
		logrus.ErrorLevel: rotateLog(fileName, "error"),
		logrus.FatalLevel: rotateLog(fileName, "fatal"),
		logrus.PanicLevel: rotateLog(fileName, "panic"),
	}
	fileMap[fileName] = writeMap
	return writeMap
}

func rotateLog(path, level string) *rotatelogs.RotateLogs {
	data, _ := rotatelogs.New(path+"-"+level+".log.%Y%m%d", rotatelogs.WithLinkName(path+"-"+level+".log"), rotatelogs.WithMaxAge(30*24*time.Hour), rotatelogs.WithRotationTime(24*time.Hour))
	return data
}

type StandardFormatter struct {
}

func (m *StandardFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	entry.Caller = getCaller()

	var fields []string
	for k, v := range entry.Data {
		fields = append(fields, fmt.Sprintf("%s=%s", k, v))
	}

	level := entry.Level
	timestamp := entry.Time.Format("2006-01-02 15:04:05")
	funPath := fmt.Sprintf("%s:%d", entry.Caller.File, entry.Caller.Line)
	fieldsStr := strings.Join(fields, " ")
	var newLog string
	switch level {
	case logrus.DebugLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s [\x1b[%dm%s\x1b[0m]\n", white, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, blue, fieldsStr)
	case logrus.InfoLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s [\x1b[%dm%s\x1b[0m]\n", green, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, blue, fieldsStr)
	case logrus.WarnLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s [\x1b[%dm%s\x1b[0m]\n", yellow, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, blue, fieldsStr)
	case logrus.ErrorLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s [\x1b[%dm%s\x1b[0m]\n", red, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, blue, fieldsStr)
	case logrus.FatalLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s [\x1b[%dm%s\x1b[0m]\n", purple, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, blue, fieldsStr)
	case logrus.PanicLevel:
		newLog = fmt.Sprintf("\x1b[%dm%s\t\x1b[0m%s \x1b[%dm%s\x1b[0m %s [\x1b[%dm%s\x1b[0m]\n", blue, strings.ToUpper(entry.Level.String()), timestamp, black, funPath, entry.Message, blue, fieldsStr)
	}
	b.WriteString(newLog)

	return b.Bytes(), nil
}

var (
	// qualified package name, cached at first use
	logrusPackage string

	// Positions in the call stack when tracing to report the calling method
	minimumCallerDepth int

	// Used for caller information initialisation
	callerInitOnce sync.Once
)

func getCaller() *runtime.Frame {
	// cache this package's fully-qualified name
	callerInitOnce.Do(func() {
		pcs := make([]uintptr, maximumCallerDepth)
		_ = runtime.Callers(0, pcs)

		// dynamic get the package name and the minimum caller depth
		for i := 0; i < maximumCallerDepth; i++ {
			funcName := runtime.FuncForPC(pcs[i]).Name()
			if strings.Contains(funcName, "getCaller") {
				logrusPackage = getPackageName(funcName)
				break
			}
		}

		minimumCallerDepth = knownLogrusFrames
	})

	// Restrict the lookback frames to avoid runaway lookups
	pcs := make([]uintptr, maximumCallerDepth)
	depth := runtime.Callers(minimumCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])

	for f, again := frames.Next(); again; f, again = frames.Next() {
		pkg := getPackageName(f.Function)

		//If the caller isn't part of this package, we're done
		if pkg == logrusPackage {
			return &f //nolint:scopelint
		}
	}

	// if we got here, we failed to find the caller's context
	return nil
}

// getPackageName reduces a fully qualified function name to the package name
// There really ought to be to be a better way...
func getPackageName(f string) string {
	for {
		lastPeriod := strings.LastIndex(f, ".")
		lastSlash := strings.LastIndex(f, "/")
		if lastPeriod > lastSlash {
			f = f[:lastPeriod]
		} else {
			break
		}
	}

	return f
}
