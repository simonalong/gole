package test

import (
	"fmt"
	"github.com/magiconair/properties"
	"github.com/simonalong/tools/util"
	"github.com/simonalong/tools/yaml"
	"io/ioutil"
	"log"
	"reflect"
	"strings"
	"testing"
)

func TestMapToProperties1(t *testing.T) {
	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = 13
	dataMap["c"] = 14

	act, err := yaml.MapToProperties(dataMap)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}
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
	//act, err := tools.MapToProperties(dataMap)
	//if err != nil {
	//	log.Fatalf("转换：%v", err)
	//	return
	//}
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
	//act, err := tools.MapToProperties(dataMap)
	//if err != nil {
	//	log.Fatalf("转换：%v", err)
	//	return
	//}
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
	//act, err := tools.MapToProperties(dataMap)
	//if err != nil {
	//	log.Fatalf("转换：%v", err)
	//	return
	//}
	//expect := "a=12\nb=13\nc=14\nd.a=inner1\nd.b=inner2\nd.c=inner3\nd.d[0]=a\nd.d[1]=b"
	//Equal(t, act, expect)
}

type ValueInnerEntity1 struct {
	Name string
	Age  int
}

func TestObjectToYaml1(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	var targetObj ValueInnerEntity1
	util.MapToObject(inner1, &targetObj)
	// {zhou 321}
	fmt.Println(util.ToJsonString(targetObj))
}

type ValueInnerEntity2 struct {
	Name   string
	Age    int
	Inner1 ValueInnerEntity1
}

func TestObjectToYaml2(t *testing.T) {
	inner1 := map[string]interface{}{}
	inner1["name"] = "inner_1"
	inner1["age"] = 1

	inner2 := map[string]interface{}{}
	inner2["name"] = "inner_2"
	inner2["age"] = 2
	inner2["inner1"] = inner1

	var targetObj ValueInnerEntity2
	util.MapToObject(inner2, &targetObj)
	// {zhou 321}
	fmt.Println(util.ToJsonString(targetObj))
}

type ValueInnerEntity3 struct {
	Name   string
	Age    int
	Inner2 ValueInnerEntity2
}

func TestObjectToYaml3(t *testing.T) {
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
	util.MapToObject(inner3, &targetObj)
	// {zhou 321}
	fmt.Println(util.ToJsonString(targetObj))
}

type ValueInnerEntity4 struct {
	Name    string
	Age     int
	DataMap map[string]ValueInnerEntity1
}

func TestObjectToYaml4(t *testing.T) {
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

	var targetObj ValueInnerEntity4
	util.MapToObject(inner1, &targetObj)
	// {zhou 321}
	fmt.Println(util.ToJsonString(targetObj))
}

func TestOb(t *testing.T) {
	//dataMap := map[string]interface{}{}
	//dataMap["a"] = 23
	//
	//mapValue := reflect.ValueOf(dataMap)
	//mV := mapValue.MapIndex(reflect.ValueOf("b"))
	//if !mV.IsValid() {
	//	fmt.Println("nil")
	//} else {
	//	fmt.Println(mV.Interface())
	//}
	//
	var mV = reflect.ValueOf(nil)
	if !mV.IsValid() {
		fmt.Println("nil")
	} else {
		fmt.Println(mV.Interface())
	}
}

func TestYamlToMap(t *testing.T) {
	yamlToMapTest(t, "./resources/yml/base.yml")
	yamlToMapTest(t, "./resources/yml/base1.yml")
	yamlToMapTest(t, "./resources/yml/array1.yml")
	yamlToMapTest(t, "./resources/yml/array2.yml")
	yamlToMapTest(t, "./resources/yml/array3.yml")
	yamlToMapTest(t, "./resources/yml/array4.yml")
	yamlToMapTest(t, "./resources/yml/array5.yml")
	//yamlToMapTest(t, "./resources/yml/array6.yml")
	yamlToMapTest(t, "./resources/yml/array7.yml")
	//yamlToMapTest(t, "./resources/yml/cron.yml")
	yamlToMapTest(t, "./resources/yml/multi_line.yml")
}

func TestPropertiesToYaml1(t *testing.T) {
	propertiesToYamlTest(t, "./resources/properties/base.properties")
	//propertiesToYamlTest(t, "./resources/properties/base1.properties")
	//propertiesToYamlTest(t, "./resources/properties/base2.properties")
	//propertiesToYamlTest(t, "./resources/properties/array1.properties")
	//propertiesToYamlTest(t, "./resources/properties/array2.properties")
	propertiesToYamlTest(t, "./resources/properties/array3.properties")
	propertiesToYamlTest(t, "./resources/properties/array4.properties")
	//propertiesToYamlTest(t, "./resources/properties/array5.properties")
	propertiesToYamlTest(t, "./resources/properties/array6.properties")
	//propertiesToYamlTest(t, "./resources/properties/array7.properties")
}

func TestYamlToKvList1(t *testing.T) {
	yamlToKvListTest(t, "./resources/yml/base.yml")
	yamlToKvListTest(t, "./resources/yml/base1.yml")
	yamlToKvListTest(t, "./resources/yml/base2.yml")
	yamlToKvListTest(t, "./resources/yml/array1.yml")
	yamlToKvListTest(t, "./resources/yml/array2.yml")
	yamlToKvListTest(t, "./resources/yml/array3.yml")
	yamlToKvListTest(t, "./resources/yml/array4.yml")
	yamlToKvListTest(t, "./resources/yml/array5.yml")
	yamlToKvListTest(t, "./resources/yml/array6.yml")
	yamlToKvListTest(t, "./resources/yml/array7.yml")
}

func TestYamlToPropertiesWithKey(t *testing.T) {
	yamlToPropertiesWithKeyTest(t, "./resources/yml/base.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/base1.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/base2.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/array1.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/array2.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/array3.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/array4.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/array5.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/array6.yml")
	yamlToPropertiesWithKeyTest(t, "./resources/yml/array7.yml")
}

func TestPropertiesToMap5(t *testing.T) {
	propertiesToMap(t, "./resources/properties/base.properties")
	propertiesToMap(t, "./resources/properties/base1.properties")
	propertiesToMap(t, "./resources/properties/base2.properties")
	propertiesToMap(t, "./resources/properties/array1.properties")
}

func propertiesToMap(t *testing.T, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		Err(t, err)
		return
	}

	expect := strings.TrimSpace(string(bytes))
	actMap, err := yaml.PropertiesToMap(expect)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}

	fmt.Println(actMap)
}

func yamlToPropertiesWithKeyTest(t *testing.T, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		Err(t, err)
		return
	}

	expect := strings.TrimSpace(string(bytes))
	property, err := yaml.YamlToPropertiesWithKey("t", expect)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}

	fmt.Println(property)
}

func yamlToKvListTest(t *testing.T, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		Err(t, err)
		return
	}

	expect := strings.TrimSpace(string(bytes))
	kvPairs, err := yaml.YamlToKvList(expect)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}

	// 获取实际数据
	actMap := map[string]string{}
	for _, pair := range kvPairs {
		actMap[pair.Left] = pair.Right
	}

	// 获取标准的数据
	property, err := yaml.YamlToProperties(expect)
	pro := properties.NewProperties()
	err = pro.Load([]byte(property), properties.UTF8)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}
	resultMap := pro.Map()

	// 数据进行对比
	for key := range resultMap {
		actValue, exist := actMap[key]
		if !exist || actValue != resultMap[key] {
			t.Errorf("有数据不一致，\n期望：key=%v, value=%v\n实际：key=%v, value=%v\n", key, resultMap[key], key, actMap[key])
		}
	}
}

func yamlToMapTest(t *testing.T, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		Err(t, err)
		return
	}
	expect := strings.TrimSpace(string(bytes))
	dataMap, err := yaml.YamlToMap(expect)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}

	value, _ := yaml.ObjectToYaml(dataMap)
	act := strings.TrimSpace(value)
	Equal(t, act, expect)
}

func propertiesToYamlTest(t *testing.T, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		Err(t, err)
		return
	}
	expect := strings.TrimSpace(string(bytes))
	yamlContent, err := yaml.PropertiesToYaml(expect)
	//fmt.Println(yamlContent)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}

	act, err := yaml.YamlToProperties(yamlContent)
	act = strings.TrimSpace(act)
	Equal(t, act, expect)
}

type TestEntity2 struct {
	Name string
	Age  int
}
