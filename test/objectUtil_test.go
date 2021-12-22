package test

import (
	"fmt"
	"github.com/simonalong/tools/util"
	"testing"
)

type ValueInnerEntityTem struct {
	Name string
	Age  ConfigValueTypeEnum
}

func TestMapToObjectTem(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = "1"

	var targetObj ValueInnerEntity1
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, util.ToJsonString(targetObj), "{\"Name\":\"inner_1\",\"Age\":1}")
}

type ValueInnerEntity1 struct {
	Name string
	Age  int
}

func TestMapToObject1(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	var targetObj ValueInnerEntity1
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1}", util.ToJsonString(targetObj))
}

type ValueInnerEntity2 struct {
	Name   string
	Age    int
	Inner1 ValueInnerEntity1
}

func TestMapToObject2(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	inner2 := map[string]interface{}{}
	inner2["name"] = "inner_2"
	inner2["age"] = 2
	inner2["inner1"] = inner1

	var targetObj ValueInnerEntity2
	_ = util.MapToObject(inner2, &targetObj)
	Equal(t, "{\"Name\":\"inner_2\",\"Age\":2,\"Inner1\":{\"Name\":\"inner_1\",\"Age\":1}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity3 struct {
	Name   string
	Age    int
	Inner2 ValueInnerEntity2
}

func TestMapToObject3(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	inner2 := map[string]interface{}{}
	inner2["name"] = "inner_2"
	inner2["age"] = 2
	inner2["inner1"] = inner1

	inner3 := map[string]interface{}{}
	inner3["name"] = "inner_3"
	inner3["age"] = 3
	inner3["inner2"] = inner2

	var targetObj ValueInnerEntity3
	_ = util.MapToObject(inner3, &targetObj)
	Equal(t, "{\"Name\":\"inner_3\",\"Age\":3,\"Inner2\":{\"Name\":\"inner_2\",\"Age\":2,\"Inner1\":{\"Name\":\"inner_1\",\"Age\":1}}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity4 struct {
	Name    string
	Age     int
	DataMap map[string]string
}

func TestMapToObject4(t *testing.T) {
	kvMap := map[string]interface{}{}
	kvMap["k1"] = "name1"
	kvMap["k2"] = "name2"

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity4
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":\"name1\",\"k2\":\"name2\"}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity5 struct {
	Name    string
	Age     int
	DataMap map[string]ValueInnerEntity1
}

func TestMapToObject5(t *testing.T) {
	v1 := map[string]interface{}{}
	v1["name"] = "inner_1"
	v1["age"] = 1

	v2 := map[string]interface{}{}
	v2["name"] = "inner_2"
	v2["age"] = 2

	kvMap := map[string]interface{}{}
	kvMap["k1"] = v1
	kvMap["k2"] = v2

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity5
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":{\"Name\":\"inner_1\",\"Age\":1},\"k2\":{\"Name\":\"inner_2\",\"Age\":2}}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity6 struct {
	Name    string
	Age     int
	DataMap map[string][]int
}

func TestMapToObject6(t *testing.T) {
	var dataList []int
	dataList = append(dataList, 12)
	dataList = append(dataList, 13)

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity6
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":[12,13],\"k2\":[12,13]}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity7 struct {
	Name    string
	Age     int
	DataMap map[string][]ValueInnerEntity1
}

func TestMapToObject7(t *testing.T) {
	var dataList []ValueInnerEntity1
	dataList = append(dataList, ValueInnerEntity1{Name: "name1", Age: 1})
	dataList = append(dataList, ValueInnerEntity1{Name: "name2", Age: 2})

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity7
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}],\"k2\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}]}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity1Tem struct {
	Name    string
	Address string
}

type ValueInnerEntity8 struct {
	Name    string
	Age     int
	DataMap map[string][]ValueInnerEntity1Tem
}

func TestMapToObject8(t *testing.T) {
	var dataList []ValueInnerEntity1
	dataList = append(dataList, ValueInnerEntity1{Name: "name1", Age: 1})
	dataList = append(dataList, ValueInnerEntity1{Name: "name2", Age: 2})

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity8
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":[{\"Name\":\"name1\",\"Address\":\"\"},{\"Name\":\"name2\",\"Address\":\"\"}],\"k2\":[{\"Name\":\"name1\",\"Address\":\"\"},{\"Name\":\"name2\",\"Address\":\"\"}]}}", util.ToJsonString(targetObj))
}

type ValueInnerEntity9Tem struct {
	Name string
	Age  string
}

type ValueInnerEntity9 struct {
	Name    string
	Age     int
	DataMap map[string][]ValueInnerEntity1
}

func TestMapToObject9(t *testing.T) {
	var dataList []ValueInnerEntity9Tem
	dataList = append(dataList, ValueInnerEntity9Tem{Name: "name1", Age: "1"})
	dataList = append(dataList, ValueInnerEntity9Tem{Name: "name2", Age: "2"})

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity9
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}],\"k2\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}]}}", util.ToJsonString(targetObj))
}

type ConfigValueTypeEnum int

const (
	YAML       ConfigValueTypeEnum = 0
	PROPERTIES ConfigValueTypeEnum = 1
	JSON       ConfigValueTypeEnum = 2
	STRING     ConfigValueTypeEnum = 3
)

type ValueInnerEntity10 struct {
	Name    string
	Age     ConfigValueTypeEnum
	DataMap map[string][]ValueInnerEntity1
}

func TestMapToObject10(t *testing.T) {
	var dataList []ValueInnerEntity9Tem
	dataList = append(dataList, ValueInnerEntity9Tem{Name: "name1", Age: "1"})
	dataList = append(dataList, ValueInnerEntity9Tem{Name: "name2", Age: "2"})

	kvMap := map[string]interface{}{}
	kvMap["k1"] = dataList
	kvMap["k2"] = dataList

	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1
	inner1["dataMap"] = kvMap

	var targetObj ValueInnerEntity9
	_ = util.MapToObject(inner1, &targetObj)
	Equal(t, "{\"Name\":\"inner_1\",\"Age\":1,\"DataMap\":{\"k1\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}],\"k2\":[{\"Name\":\"name1\",\"Age\":1},{\"Name\":\"name2\",\"Age\":2}]}}", util.ToJsonString(targetObj))
}

type ValueObjectTest1 struct {
	AppName string
	Age     int
}

func TestObjectToJson1(t *testing.T) {
	entity := ValueObjectTest1{AppName: "zhou", Age: 12}
	Equal(t, util.ObjectToJson(entity), "{\"age\":12,\"appName\":\"zhou\"}")
}

type ValueObjectTest2 struct {
	AppName string

	Age1 int
	Age2 int8
	Age3 int16
	Age4 int32
	Age5 int64

	UAge1 uint
	UAge2 uint8
	UAge3 uint16
	UAge4 uint32
	UAge5 uint64

	FAge1 float32
	FAge2 float64

	CAge1 complex64
	CAge2 complex128
}

func TestObjectToJson2(t *testing.T) {
	entity := ValueObjectTest2{
		AppName: "zhou",
		Age1:    12,
		Age2:    12,
		Age3:    12,
		Age4:    12,
		Age5:    12,
		UAge1:   12,
		UAge2:   12,
		UAge3:   12,
		UAge4:   12,
		UAge5:   12,
		FAge1:   12.1,
		FAge2:   12.2,
		CAge1:   3.2 + 12i,
		CAge2:   5.2 + 13i,
	}
	Equal(t, util.ObjectToJson(entity), "{\"age1\":12,\"age2\":12,\"age3\":12,\"age4\":12,\"age5\":12,\"appName\":\"zhou\",\"cAge1\":\"(3.2+12i)\",\"cAge2\":\"(5.2+13i)\",\"fAge1\":12.1,\"fAge2\":12.2,\"uAge1\":12,\"uAge2\":12,\"uAge3\":12,\"uAge4\":12,\"uAge5\":12}")
}

type ValueObjectTest3 struct {
	AppName []string
	Age1    map[string]interface{}
}

func TestObjectToJson3(t *testing.T) {
	var arrays []string
	arrays = append(arrays, "zhou")
	arrays = append(arrays, "wang")

	dataMap := map[string]interface{}{}
	dataMap["a"] = 1
	dataMap["b"] = 2

	entity := ValueObjectTest3{
		AppName: arrays,
		Age1:    dataMap,
	}
	Equal(t, util.ObjectToJson(entity), "{\"age1\":{\"a\":1,\"b\":2},\"appName\":[\"zhou\",\"wang\"]}")
}

type ValueObjectTest4 struct {
	AppName string
	Inner   ValueObjectTest3
}

func TestObjectToJson4(t *testing.T) {
	var arrays []string
	arrays = append(arrays, "zhou")
	arrays = append(arrays, "wang")

	dataMap := map[string]interface{}{}
	dataMap["a"] = 1
	dataMap["b"] = 2

	entity3 := ValueObjectTest3{
		AppName: arrays,
		Age1:    dataMap,
	}

	var entity4 ValueObjectTest4
	entity4.Inner = entity3
	entity4.AppName = "zhou"
	fmt.Println(util.ObjectToJson(entity4))
	Equal(t, util.ObjectToJson(entity4), "{\"appName\":\"zhou\",\"inner\":{\"age1\":{\"a\":1,\"b\":2},\"appName\":[\"zhou\",\"wang\"]}}")
}
