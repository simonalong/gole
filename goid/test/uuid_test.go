package test

import (
	"fmt"
	"github.com/simonalong/gole/goid"
	"testing"
)

func TestUUID(t *testing.T) {
	id := goid.GenerateUUID()
	// 9073dde4-f7ab-456f-8ce1-3d3613963159
	fmt.Println(id)
}

func TestUUIDFull(t *testing.T) {
	fmt.Println(goid.GenerateUUIDFullString())
	fmt.Println(goid.GenerateUUIDFullString())
	fmt.Println(goid.GenerateUUIDFullString())
	fmt.Println(goid.GenerateUUIDFullString()[:10])
}
