package test

import (
	"testing"

	"github.com/simonalong/gole/util"
)

type AssociateStruct struct {
	Key  string
	Name string
	Age  int
}

func initList() []AssociateStruct {
	return []AssociateStruct{
		{"K", "库陈胜", 20},
		{"K", "库陈胜", 30},
		{"K1", "库陈胜", 30},
		{"K1", "库陈胜", 40},
		{"K1", "库陈胜", 50},
		{"K2", "库陈胜", 60},
		{"K2", "库陈胜", 70},
		{"K2", "库陈胜", 80},
		{"K4", "库陈胜", 90},
	}
}

var transformFun = func(a AssociateStruct) util.Pair[string, AssociateStruct] {
	return util.NewPair(a.Key, a)
}

var transformFun1 = func(a AssociateStruct) int {
	return a.Age
}

var keySelector = func(a AssociateStruct) string {
	return a.Key
}

func TestAssociate(t *testing.T) {
	list := initList()
	l := util.Associate(list, transformFun)
	t.Logf("%v", l)
}

func TestAssociateTo(t *testing.T) {
	list := initList()
	m := map[string]AssociateStruct{}
	r := util.AssociateTo(list, &m, transformFun)
	t.Logf("%v", r)
}

func TestAssociateBy(t *testing.T) {
	list := initList()
	r := util.AssociateBy(list, keySelector)
	t.Logf("%v", r)
}

func TestAssociateByAndValue(t *testing.T) {
	list := initList()
	r := util.AssociateByAndValue(list, keySelector, transformFun)
	t.Logf("%v", r)
}

func TestAssociateByAndValueTo(t *testing.T) {
	list := initList()
	m := make(map[string]int)
	util.AssociateByAndValueTo(list, &m, keySelector, transformFun1)
	t.Logf("%v", m)
}

//
func TestAssociateByTo(t *testing.T) {
	list := initList()
	m := make(map[string]AssociateStruct)
	util.AssociateByTo(list, &m, keySelector)
	t.Logf("%v", m)
}

//

//
func TestAssociateWith(t *testing.T) {
	list := initList()
	m := util.AssociateWith(list, transformFun1)
	t.Logf("%v", m)
}

//
func TestAssociateWithTo(t *testing.T) {
	list := initList()
	m := make(map[AssociateStruct]int)
	util.AssociateWithTo(list, &m, transformFun1)
	t.Logf("%v", m)
}
