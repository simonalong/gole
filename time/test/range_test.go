package test

import (
	"github.com/magiconair/properties/assert"
	"github.com/simonalong/gole/time"
	"testing"
	t0 "time"
)

func TestRange1(t *testing.T) {
	var t1, t2 t0.Time
	t1, _ = time.ParseTime("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.122")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1毫秒")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:01.123")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1秒")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:31:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1分钟")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 11:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1小时")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-11 12:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1天")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-06-12 12:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1月")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2023-07-12 12:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1年 6天")
}

func TestRange2(t *testing.T) {
	var t1, t2 t0.Time
	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 11:12:10.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.122")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1小时 19分钟 49秒 999毫秒")
}

func TestRangeEn1(t *testing.T) {
	var t1, t2 t0.Time
	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.122")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1ms")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:01.123")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1s")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:31:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1m")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 11:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1h")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-11 12:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1d")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-06-12 12:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1M")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.123")
	t2, _ = time.ParseTimeYmdHmsS("2023-07-12 12:32:00.123")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1y 6d")
}

func TestRangeEn2(t *testing.T) {
	var t1, t2 t0.Time
	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 11:12:10.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.122")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1h 19m 49s 999ms")

	t1, _ = time.ParseTimeYmdHmsS("2024-07-12 11:12:10.123")
	t2, _ = time.ParseTimeYmdHmsS("2024-07-12 12:32:00.122")
	assert.Equal(t, time.ParseDurationOfTimeSecondForViewEn(t1, t2), "1h 19m 49s")
}

func TestRangeUs(t *testing.T) {
	var t1, t2 t0.Time
	t1, _ = time.ParseTimeYmdHmsSuS("2024-07-12 11:12:10.123333")
	t2, _ = time.ParseTimeYmdHmsSuS("2024-07-12 11:12:10.123332")
	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1us")
	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1微秒")
}

// 这个也不是很准确，还是存在差距，因此这里注释
//func TestRange(t *testing.T) {
//	var t1, t2 t0.Time
//	t1 = t0.Now()
//	t0.Sleep(1 * t0.Second)
//	t2 = t0.Now()
//	assert.Equal(t, time.ParseDurationOfTimeForViewEn(t1, t2), "1s")
//	assert.Equal(t, time.ParseDurationOfTimeForView(t1, t2), "1秒")
//}

// 这个不是很准确，有时候会差距1毫秒，这里就注释
//func TestRange3(t *testing.T) {
//	var t1 t0.Time
//	t1 = t0.Now()
//	t0.Sleep(1 * t0.Second)
//	assert.Equal(t, time.ParseDurationForView(t0.Now().Sub(t1)), "1秒")
//}
