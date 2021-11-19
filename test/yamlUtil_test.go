package test

import (
	"github.com/simonalong/tools"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

func TestMapToProperties1(t *testing.T) {
	dataMap := map[string]interface{}{}
	dataMap["a"] = 12
	dataMap["b"] = 13
	dataMap["c"] = 14

	act, err := tools.MapToProperties(dataMap)
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
	//propertiesToYamlTest(t, "./resources/properties/base.properties")
	propertiesToYamlTest(t, "./resources/properties/base1.properties")
}

func yamlToMapTest(t *testing.T, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		Err(t, err)
		return
	}
	expect := strings.TrimSpace(string(bytes))
	dataMap, err := tools.YamlToMap(expect)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}

	act := strings.TrimSpace(tools.MapToYaml(dataMap))
	//fmt.Println(act)
	//fmt.Println(expect)
	Equal(t, act, expect)
}

func propertiesToYamlTest(t *testing.T, filePath string) {
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		Err(t, err)
		return
	}
	expect := strings.TrimSpace(string(bytes))
	yamlContent, err := tools.PropertiesToYaml(expect)
	if err != nil {
		log.Fatalf("转换错误：%v", err)
		return
	}

	act, err := tools.YamlToProperties(yamlContent)
	act = strings.TrimSpace(act)
	//fmt.Println(act)
	//fmt.Println(expect)
	Equal(t, act, expect)
}
