package time

import (
	"fmt"
	"strings"
	t0 "time"
)

const (
	//micros	我们这里最小的单位是微秒
	micros = 1
	millis = 1000 * micros
	second = 1000 * millis
	minute = 60 * second
	hour   = 60 * minute
	day    = 24 * hour
	month  = 30 * day
	year   = 12 * month
)

// ParseDurationOfTimeForView 计算两个时间之间的差值，主要以可视化展示；默认中文
// 示例：1年 2月 4天 1分钟 12秒 132毫秒
func ParseDurationOfTimeForView(time1, time2 t0.Time) string {
	if time1.Before(time2) {
		return ParseDurationOfInt64ForView(time2.Sub(time1).Microseconds(), false, false)
	} else {
		return ParseDurationOfInt64ForView(time1.Sub(time2).Microseconds(), false, false)
	}
}

// ParseDurationOfTimeForViewEn 计算两个时间之间的差值，主要以可视化展示；英文
// 示例：1h 19m 49s 999ms
func ParseDurationOfTimeForViewEn(time1, time2 t0.Time) string {
	if time1.Before(time2) {
		return ParseDurationOfInt64ForView(time2.Sub(time1).Microseconds(), false, true)
	} else {
		return ParseDurationOfInt64ForView(time1.Sub(time2).Microseconds(), false, true)
	}
}

func ParseDurationOfTimeSecondForViewEn(time1, time2 t0.Time) string {
	if time1.Before(time2) {
		return ParseDurationSecondForViewEn(time2.Sub(time1))
	} else {
		return ParseDurationSecondForViewEn(time1.Sub(time2))
	}
}

func ParseDurationOfTimeSecondForViewCn(time1, time2 t0.Time) string {
	if time1.Before(time2) {
		return ParseDurationSecondForViewCn(time2.Sub(time1))
	} else {
		return ParseDurationSecondForViewCn(time1.Sub(time2))
	}
}

func ParseDurationSecondForViewEn(duration t0.Duration) string {
	return strings.TrimSpace(ParseDurationOfInt64ForView(duration.Microseconds(), true, true))
}

func ParseDurationSecondForViewCn(duration t0.Duration) string {
	return strings.TrimSpace(ParseDurationOfInt64ForView(duration.Microseconds(), true, false))
}

func ParseDurationForView(duration t0.Duration) string {
	return strings.TrimSpace(ParseDurationOfInt64ForView(duration.Microseconds(), false, false))
}

func ParseDurationForViewEn(duration t0.Duration) string {
	return strings.TrimSpace(ParseDurationOfInt64ForView(duration.Microseconds(), false, true))
}

func ParseDurationOfInt64ForView(duration int64, isSecond, isEn bool) string {
	if isSatisfyYear(duration) {
		return strings.TrimSpace(fmt.Sprintf("%v", duration/year) + yearStr(isEn) + parseMonth(duration%year, isSecond, isEn))
	} else {
		return strings.TrimSpace(parseMonth(duration%year, isSecond, isEn))
	}
}

func parseMonth(duration int64, isSecond, isEn bool) string {
	if isSatisfyMonth(duration) {
		return fmt.Sprintf("%v", duration/month) + monthStr(isEn) + parseDay(duration%month, isSecond, isEn)
	} else {
		return parseDay(duration%month, isSecond, isEn)
	}
}

func parseDay(duration int64, isSecond, isEn bool) string {
	if isSatisfyDay(duration) {
		return fmt.Sprintf("%v", duration/day) + dayStr(isEn) + parseHour(duration%day, isSecond, isEn)
	} else {
		return parseHour(duration%day, isSecond, isEn)
	}
}

func parseHour(duration int64, second, isEn bool) string {
	if isSatisfyHour(duration) {
		return fmt.Sprintf("%v", duration/hour) + hourStr(isEn) + parseMinute(duration%hour, second, isEn)
	} else {
		return parseMinute(duration%hour, second, isEn)
	}
}

func parseMinute(duration int64, isSecond, isEn bool) string {
	if isSatisfyMinute(duration) {
		return fmt.Sprintf("%v", duration/minute) + minuteStr(isEn) + parseSecond(duration%minute, isSecond, isEn)
	} else {
		return parseSecond(duration%minute, isSecond, isEn)
	}
}

func parseSecond(duration int64, isSecond, isEn bool) string {
	if isSatisfySecond(duration) {
		if isSecond {
			return fmt.Sprintf("%v", duration/second) + secondStr(isEn)
		} else {
			return fmt.Sprintf("%v", duration/second) + secondStr(isEn) + parseMillis(duration%second, isEn)
		}
	} else {
		if isSecond {
			return "0" + secondStr(isEn)
		} else {
			return parseMillis(duration%second, isEn)
		}
	}
}

func parseMillis(duration int64, isEn bool) string {
	if isSatisfyMillis(duration) {
		return fmt.Sprintf("%v", duration/millis) + millisStr(isEn) + parseMicros(duration%millis, isEn)
	} else {
		return parseMicros(duration%millis, isEn)
	}
}

func parseMicros(duration int64, isEn bool) string {
	if isSatisfyMicro(duration) {
		return fmt.Sprintf("%v", duration/micros) + microStr(isEn)
	}
	return ""
}

func isSatisfyYear(duration int64) bool {
	return duration >= year
}

func isSatisfyMonth(duration int64) bool {
	return duration >= month
}
func isSatisfyDay(duration int64) bool {
	return duration >= day
}
func isSatisfyHour(duration int64) bool {
	return duration >= hour
}

func isSatisfyMinute(duration int64) bool {
	return duration >= minute
}

func isSatisfySecond(duration int64) bool {
	return duration >= second
}

func isSatisfyMillis(duration int64) bool {
	return duration >= millis
}

func isSatisfyMicro(duration int64) bool {
	return duration >= micros
}

func yearStr(isEn bool) string {
	if isEn {
		return "y "
	}
	return "年 "
}

func monthStr(isEn bool) string {
	if isEn {
		return "M "
	}
	return "月 "
}

func dayStr(isEn bool) string {
	if isEn {
		return "d "
	}
	return "天 "
}

func hourStr(enOrCn bool) string {
	if enOrCn {
		return "h "
	}
	return "小时 "
}

func minuteStr(enOrCn bool) string {
	if enOrCn {
		return "m "
	}
	return "分钟 "
}

func secondStr(enOrCn bool) string {
	if enOrCn {
		return "s "
	}
	return "秒 "
}

func millisStr(enOrCn bool) string {
	if enOrCn {
		return "ms "
	}
	return "毫秒 "
}

func microStr(enOrCn bool) string {
	if enOrCn {
		return "us "
	}
	return "微秒 "
}
