package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	t0 "time"

	"github.com/simonalong/gole/time"
)

func TestTime(t *testing.T) {
	t1 := time.TimeInMillis()
	t2 := time.TimeInSeconds()
	t3 := time.TimeInNano()
	t4 := time.TimeInMicro()
	t.Logf("t1: %v, t2: %v", t1, t2)
	t.Logf("t3: %v, t4: %v", t3, t4)
}

func TestTimeBetween(t *testing.T) {
	athen := t0.Date(2000, t0.January, 1, 12, 0, 0, 0, t0.UTC)
	anow := t0.Date(2022, t0.March, 1, 16, 21, 0, 0, t0.UTC)
	t.Logf("years between: %v", time.YearsBetween(anow, athen))
	t.Logf("months between: %v", time.MonthsBetween(anow, athen))
	t.Logf("days between: %v", time.DaysBetween(anow, athen))
	t.Logf("hours between: %v", time.HoursBetween(anow, athen))
	t.Logf("minutes between: %v", time.MinutesBetween(anow, athen))
	t.Logf("seconds between: %v", time.SecondsBetween(anow, athen))
	t.Logf("milliseconds between: %v", time.MilliSecondsBetween(anow, athen))
}

func TestTimeSpan(t *testing.T) {
	athen := t0.Date(2000, t0.January, 1, 12, 0, 0, 0, t0.UTC)
	anow := t0.Date(2022, t0.March, 1, 16, 21, 0, 0, t0.UTC)
	t.Logf("years between: %v", time.YearSpan(anow, athen))
	t.Logf("months between: %v", time.MonthSpan(anow, athen))
	t.Logf("days between: %v", time.DaySpan(anow, athen))
	t.Logf("hours between: %v", time.HourSpan(anow, athen))
	t.Logf("minutes between: %v", time.MinuteSpan(anow, athen))
	t.Logf("seconds between: %v", time.SecondSpan(anow, athen))
	t.Logf("milliseconds between: %v", time.MilliSecondSpan(anow, athen))
}

func TestNumToTimeDuration(t *testing.T) {
	data := time.NumToTimeDuration(3, t0.Hour)
	fmt.Println(data.Milliseconds())
}

func TestParseTime(t *testing.T) {
	d, _ := time.ParseTime("20220729")
	fmt.Println(time.TimeToStringYmdHmsS(d))
}

func TestToTime(t *testing.T) {
	milliseconds := time.Now().UnixMilli()
	fmt.Println(time.MillisecondToTime(milliseconds))
}

func TestMiddleTime(t *testing.T) {
	t1, _ := time.ParseTimeYmdHmsS("2024-08-01 12:23:00.321")
	t2, _ := time.ParseTimeYmdHmsS("2024-01-01 12:23:00.321")

	t3 := time.GetMiddleTime(t1, t2)
	fmt.Println(time.TimeToStringYmdHmsS(t3))
}

func TestRandomSleep(t *testing.T) {
	time.RandomSleep(10, t0.Second)
}

func TestAddTime(t *testing.T) {
	currentTime, _ := time.ParseTimeYmdHmsS("2024-08-01 12:23:00.321")

	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHours(currentTime, 2)), "2024-08-01 14:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "2")), "2024-08-01 14:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "2h")), "2024-08-01 14:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "+2")), "2024-08-01 14:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "+2h")), "2024-08-01 14:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "-2")), "2024-08-01 10:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "-2h")), "2024-08-01 10:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "-13")), "2024-07-31 23:23:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddHoursStr(currentTime, "-13h")), "2024-07-31 23:23:00.321")

	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutes(currentTime, 2)), "2024-08-01 12:25:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutesStr(currentTime, "2")), "2024-08-01 12:25:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutesStr(currentTime, "+2")), "2024-08-01 12:25:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutesStr(currentTime, "+2m")), "2024-08-01 12:25:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutesStr(currentTime, "-2")), "2024-08-01 12:21:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutesStr(currentTime, "-2m")), "2024-08-01 12:21:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutesStr(currentTime, "-24")), "2024-08-01 11:59:00.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMinutesStr(currentTime, "-24m")), "2024-08-01 11:59:00.321")

	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddSeconds(currentTime, 2)), "2024-08-01 12:23:02.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddSecondsStr(currentTime, "2")), "2024-08-01 12:23:02.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddSecondsStr(currentTime, "+2")), "2024-08-01 12:23:02.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddSecondsStr(currentTime, "+2s")), "2024-08-01 12:23:02.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddSecondsStr(currentTime, "-2")), "2024-08-01 12:22:58.321")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddSecondsStr(currentTime, "-2s")), "2024-08-01 12:22:58.321")

	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMilliseconds(currentTime, 2)), "2024-08-01 12:23:00.323")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMillisecondsStr(currentTime, "2")), "2024-08-01 12:23:00.323")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMillisecondsStr(currentTime, "+2")), "2024-08-01 12:23:00.323")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMillisecondsStr(currentTime, "+2ms")), "2024-08-01 12:23:00.323")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMillisecondsStr(currentTime, "-2")), "2024-08-01 12:23:00.319")
	assert.Equal(t, time.TimeToStringYmdHmsS(time.AddMillisecondsStr(currentTime, "-2ms")), "2024-08-01 12:23:00.319")
}
