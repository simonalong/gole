package test

import (
	"github.com/simonalong/tools"
	"testing"
)

func TestMapToProperties1(t *testing.T) {
	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = 13
	dataMap["c"] = 14

	act := tools.MapToProperties(dataMap)
	expect := "a=12\nb=13\nc=14\n"
	Equal(t, act, expect)
}

func TestMapToProperties2(t *testing.T) {
	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = 13
	dataMap["c"] = 14

	innerMap1 := map[string]interface{}{}
	innerMap1["a"] = "inner1"
	innerMap1["b"] = "inner2"
	innerMap1["c"] = "inner3"
	dataMap["d"] = innerMap1

	// 顺序不固定，无法测试
	//act := tools.MapToProperties(dataMap)
	//expect := "a=12\nb=13\nc=14\nd.a=inner1\nd.b=inner2\nd.c=inner3"
	//Equal(t, act, expect)
}

func TestMapToProperties3(t *testing.T) {
	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = 13
	dataMap["c"] = 14

	innerMap1 := map[string]interface{}{}
	innerMap1["a"] = "inner1"
	innerMap1["b"] = "inner2"
	innerMap1["c"] = "inner3"
	dataMap["d"] = innerMap1

	array := []string{}
	array = append(array, "a")
	array = append(array, "b")
	dataMap["e"] = array

	// 顺序不固定，无法测试
	//act := tools.MapToProperties(dataMap)
	//expect := "a=12\nb=13\nc=14\nd.a=inner1\nd.b=inner2\nd.c=inner3\ne[0]=a\ne[1]=b"
	//Equal(t, act, expect)
}

func TestMapToProperties4(t *testing.T) {
	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = 13
	dataMap["c"] = 14

	innerMap1 := map[string]interface{}{}
	innerMap1["a"] = "inner1"
	innerMap1["b"] = "inner2"
	innerMap1["c"] = "inner3"
	array := []string{}
	array = append(array, "a")
	array = append(array, "b")
	innerMap1["d"] = array
	dataMap["d"] = innerMap1

	// 顺序不固定，无法测试
	//act := tools.MapToProperties(dataMap)
	//expect := "a=12\nb=13\nc=14\nd.a=inner1\nd.b=inner2\nd.c=inner3\nd.d[0]=a\nd.d[1]=b"
	//Equal(t, act, expect)
}
