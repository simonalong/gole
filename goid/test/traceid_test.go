package test

import (
	"testing"

	"github.com/simonalong/gole/goid"
)

func TestTraceId(t *testing.T) {
	tid := goid.GenerateTraceID()
	t.Logf("trace id: %s", tid)
	tid2 := goid.GenerateTraceID()
	t.Logf("trace id: %s", tid2)
	tid3 := goid.GenerateTraceID()
	t.Logf("trace id: %s", tid3)
}
