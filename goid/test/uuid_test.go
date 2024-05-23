package test

import (
	"github.com/simonalong/gole/goid"
	"testing"
)

func TestUUID(t *testing.T) {
	id := goid.GenerateUUID()
	t.Logf("UUID: %s\n", id)
}
