package test

import (
	"testing"

	"github.com/simonalong/gole/util"
)

type sliceTestStruct struct {
	Name string
	Age  int
}

func TestSliceDistinctTo(t *testing.T) {
	s1 := sliceTestStruct{
		Name: "库陈胜",
		Age:  30,
	}
	s2 := sliceTestStruct{
		Name: "酷达舒",
		Age:  29,
	}
	s3 := sliceTestStruct{
		Name: "库陈胜",
		Age:  28,
	}
	list := []sliceTestStruct{s1, s2, s3}
	l := util.SliceDistinctTo(list, func(s sliceTestStruct) string {
		return s.Name
	})
	t.Logf("%s\n", util.ToString(l))
	b := util.SliceContains(list, func(s sliceTestStruct) string {
		return s.Name
	}, "库陈胜")
	t.Logf("%v\n", b)
}
