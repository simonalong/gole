package time

import "time"

var (
	Year  = "2006"
	Month = "01"
	Day   = "02"

	Hour   = "15"
	Minute = "04"
	Second = "05"

	yMdHmsSSS = "2006-01-02 15:04:05.000"
	yMdHmsS   = "2006-01-02 15:04:05.0"
	yMdHms    = "2006-01-02 15:04:05"
	yMdHm     = "2006-01-02 15:04"
	yMdH      = "2006-01-02 15"
	yMd       = "2006-01-02"
	yM        = "2006-01"
	y         = "2006"
	yyyyMMdd  = "20060102"

	HmsSSSMore = "15:04:05.000000000"
	HmsSSS     = "15:04:05.000"
	Hms        = "15:04:05"
	Hm         = "15:04"
	H          = "15"
)

func TimeToStringYmdHms(t time.Time) string {
	return t.Format(yMdHms)
}

func TimeToStringYmdHmsS(t time.Time) string {
	return t.Format(yMdHmsSSS)
}

func TimeToStringFormat(t time.Time, format string) string {
	return t.Format(format)
}

func ParseTimeYmsHms(timeStr string) (time.Time, error) {
	return time.ParseInLocation(Year+"-"+Month+"-"+Day+" "+Hour+":"+Minute+":"+Second, timeStr, time.Local)
}

func ParseTimeYmsHmsS(timeStr string) (time.Time, error) {
	return time.ParseInLocation(Year+"-"+Month+"-"+Day+" "+Hour+":"+Minute+":"+Second+".000", timeStr, time.Local)
}

func ParseTimeYmsHmsLoc(timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(Year+"-"+Month+"-"+Day+" "+Hour+":"+Minute+":"+Second, timeStr, loc)
}

func ParseTimeYmsHmsSLoc(timeStr string, loc *time.Location) (time.Time, error) {
	return time.ParseInLocation(Year+"-"+Month+"-"+Day+" "+Hour+":"+Minute+":"+Second+".000", timeStr, loc)
}
