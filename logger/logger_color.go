package logger

import (
	"github.com/gookit/color"
	"strings"
)

func colorTimestampStr(level, message string) string {
	level = strings.ToUpper(level)
	switch level {
	case "PANIC":
		return color.BgLightMagenta.Render(message)
	case "FATAL":
		return color.BgLightMagenta.Render(message)
	case "ERROR":
		return color.BgLightRed.Render(message)
	case "WARN", "WARNING":
		return color.BgLightYellow.Render(message)
	case "DEBUG":
		return color.BgLightGreen.Render(message)
	default:
		return color.BgDarkGray.Render(message)
	}
}

func colorMsgStr(level, message string) string {
	level = strings.ToUpper(level)
	switch level {
	case "PANIC":
		return color.BgLightMagenta.Render(message)
	case "FATAL":
		return color.BgLightMagenta.Render(message)
	case "ERROR":
		return color.Red.Render(message)
	case "WARN", "WARNING":
		return color.Yellow.Render(message)
	case "INFO":
		return color.FgLightBlue.Render(message)
	case "DEBUG":
		return color.FgLightGreen.Render(message)
	case "TRACE":
		return color.FgLightCyan.Render(message)
	}
	return message
}

func colorLevelStr(level string) string {
	level = strings.ToUpper(level)
	switch level {
	case "PANIC":
		return color.BgLightMagenta.Render(level)
	case "FATAL":
		return color.BgLightMagenta.Render(level)
	case "ERROR":
		return color.BgLightRed.Render(level)
	case "WARN", "WARNING":
		return color.BgLightYellow.Render(level)
	case "INFO":
		return color.FgLightBlue.Render(level)
	case "DEBUG":
		return color.FgLightGreen.Render(level)
	case "TRACE":
		return color.FgLightCyan.Render(level)
	}
	return level
}
