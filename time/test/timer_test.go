package test

import (
	"fmt"
	"testing"

	t0 "time"

	"github.com/simonalong/gole/time"
)

func TestTimer(t *testing.T) {

	// 这个定时器的打印，执行后是不会立马执行的，要等5秒之后才会执行
	timer := time.NewTimerWithFire(5, func(tm *time.Timer) {
		fmt.Println("打印一次", time.TimeToStringYmdHms(time.Now()))
	})
	timer.Start()
	t0.Sleep(6 * t0.Second)
}

func TestTimer1(t *testing.T) {

	globalCount := 0
	timer := time.NewTimerWithFire(11, func(tm *time.Timer) {
		// 这里是运行在协程内的
		globalCount++
		t.Logf("globalCount: %d", globalCount)
	})
	timer.Start()
	t.Logf("timer.stopped = %v\n", timer.IsStopped)

	// 等待timer执行5次
	for globalCount != 5 {
		t0.Sleep(t0.Second)
	}

	timer.Stop()
	t.Logf("timer.stopped = %v\n", timer.IsStopped)

	// 证明timer已停
	for globalCount != 10 {
		t0.Sleep(t0.Second)
		globalCount++
	}
}

func TestTimerParam(t *testing.T) {
	globalCount := 0
	timer := time.NewTimerWithFire(14, func(tm *time.Timer) {
		// 这里是运行在协程内的
		globalCount++
		t.Logf("globalCount: %d", globalCount)
	})
	timer.Start()
	t.Logf("timer.stopped = %v\n", timer.IsStopped)
	// 等待timer执行3次
	for globalCount != 3 {
		t0.Sleep(t0.Second)
	}
	// 修改参数
	timer.SetInterval(3)
	timer.SetOnTimer(func(tm *time.Timer) {
		globalCount++
		t.Logf("globalCount2: %d", globalCount)
	})
	// 等待timer执行3次
	for globalCount != 6 {
		t0.Sleep(t0.Second)
	}
	timer.Stop()
}
