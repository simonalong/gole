package test

import (
	"testing"

	"github.com/simonalong/gole/util"
)

func TestIndexOf(t *testing.T) {
	list := []int{2, 4, 6, 9, 12}
	item := 10
	res := isc.IndexOf(list, item)
	t.Logf("%v\n", res)
}
