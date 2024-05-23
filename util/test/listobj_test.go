package test

import (
	"testing"

	"github.com/simonalong/gole/util"
)

type MyStruct struct {
	Name string
	Age  int
}

func TestISCList_associateBy(t *testing.T) {
	var testList util.ISCList[MyStruct]
	s1 := MyStruct{
		Name: "K",
		Age:  1,
	}
	testList.Add(s1)
	testList.Add(MyStruct{
		Name: "K2",
		Age:  2,
	})
	testList.Add(MyStruct{
		Name: "K3",
		Age:  3,
	})
	//l := util.AssociateBy(testList, func(t MyStruct) any {
	//	return t.Name
	//})
	//t.Logf("%v\n", util.ToString(l))
}

func TestNewListWithList(t *testing.T) {
	//list := []string{"1","2"}
	l := util.NewList[string]()
	t.Logf("%v", l)
}

func TestISCList_Add(t *testing.T) {
	l := util.NewList[string]()
	l.Add("3")
	t.Logf("%v", l)
}

func TestISCList_AddAll(t *testing.T) {
	l := util.NewList[string]()
	l.AddAll("4", "5", "6")
	t.Logf("%v", l)
}

func TestISCList_Insert(t *testing.T) {
	l := util.NewListWithItems("1", "2", "3")
	l.Insert(2, "7")
	t.Logf("%v", l)
}

func TestISCList_Delete(t *testing.T) {
	l := util.NewListWithItems("1", "2", "3", "4")
	l.Delete(2)
	t.Logf("%v", l)
}

func TestISCList_Clear(t *testing.T) {
	l := util.NewListWithItems("1", "2", "3", "4")
	l.Clear()
	t.Logf("%v", l)
}

func TestISCList_IsEmpty(t *testing.T) {
	l := util.NewListWithItems("1", "2", "3", "4")
	r := l.IsEmpty()
	t.Logf("%v", r)
	l.Clear()
	r = l.IsEmpty()
	t.Logf("%v", r)
}

func TestISCList_ForEach(t *testing.T) {
	l := util.NewListWithItems("1", "2", "3", "4")
	l.ForEach(func(item string) {
		t.Logf("%s", item)
	})
}
