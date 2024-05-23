package test

import (
	"github.com/simonalong/gole/util"
	"testing"
)

func TestSet(t *testing.T) {
	l := util.NewListWithItems(1, 2, 3, 4, 1, 2, 3, 4, 5, 6)
	t.Logf("%v", l)
	ls := util.ListToSet(l)
	t.Logf("%v", ls)

	s := util.NewSetWithItems(1, 2, 3, 4, 1, 2, 3, 4, 5, 6)
	t.Logf("%v", s)

	_ = s.Add(7)
	s.AddAll(8, 9)
	t.Logf("%v", s)
	_ = s.Delete(5)
	t.Logf("%v", s)
	s.Clear()
	t.Logf("%v", s)
}
