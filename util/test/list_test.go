package test

import (
	"fmt"
	"github.com/simonalong/gole/util"
	"testing"
)

func TestList(t *testing.T) {
	list := util.NewList[string]()
	list.Add("aaaa")
	list.AddAll("bbbb", "ccc")

	t.Logf("list: %v\n", list)

	list.Clear()
	t.Logf("list: %v\n", list)

}

func TestForeach(t *testing.T) {
	list := util.NewList[string]()
	list.AddAll("aaaa", "bbbb", "ccc")
	for i, s := range list {
		fmt.Println(i, s)
	}
	list.ForEachIndexed(func(i int, s string) {
		fmt.Println(i, s)
	})
}

func TestMap(t *testing.T) {
	m := util.NewMap[string, string]()
	m["aa"] = "bb"
	m.Put("cc", "dd")
	t.Logf("m: %v\n", m)
	m.Delete("cc")
	t.Logf("m: %v\n", m)
	m.Clear()
	t.Logf("m: %v\n", m)
}
