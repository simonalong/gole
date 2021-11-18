package tools

import (
	"gopkg.in/yaml.v2"
	"log"
	"reflect"
	"strconv"
	"strings"
)

/**
 *  1.yaml <---> properties
 *  2.yaml <---> json
 *  3.yaml <---> map
 *  4.yaml <---> list
 *  5.yaml <---> kvList
 */

type KeyValue struct {
	Key   string
	Value interface{}
}

//func YamlToPropertiesStr(contentOfYaml string) string {
//	// yaml 到 map
//	dataMap := YamlToMap(contentOfYaml)
//
//}
//
//
//
//func PropertiesStrToYaml(contentOfProperties string) string {
//
//}
//
func YamlToMap(contentOfYaml string) map[string]interface{} {
	resultMap := make(map[string]interface{})
	err := yaml.Unmarshal([]byte(contentOfYaml), &resultMap)
	if err != nil {
		log.Fatalf("YamlToMap error: %v", err)
		return nil
	}

	return resultMap
}

func MapToYaml(dataMap map[string]interface{}) string {
	bytes2, err := yaml.Marshal(dataMap)
	if err != nil {
		log.Fatalf("MapToYaml error: %v", err)
		return ""
	}
	return string(bytes2)
}

//
//func YamlToJson(contentOfYaml string) string {
//
//}
//
//func JsonToYaml(contentOfJson string) string {
//
//}
//
//func YamlToList(contentOfYaml string) []interface{} {
//
//}
//
//func KvListToYaml(kvList []KeyValue) string {
//
//}

// 进行深层嵌套的map数据处理
func MapToProperties(dataMap map[string]interface{}) {
	propertyStrList := []string{}
	for key, value := range dataMap {
		valueKind := reflect.TypeOf(value).Kind()
		switch valueKind {
		case reflect.Map:
			{
				propertyStrList = append(propertyStrList, doMapToProperties(propertyStrList, value, prefixWithDOT("")+key)...)
			}
		case reflect.Array, reflect.Slice:
			{
				objectValue := reflect.ValueOf(value)
				for index := 0; index < objectValue.Len(); index++ {
					propertyStrList = append(propertyStrList, doMapToProperties(propertyStrList, objectValue.Index(index), prefixWithDOT("")+key+"["+strconv.Itoa(index)+"]")...)
				}
			}
		case reflect.String:
			objectValue := reflect.ValueOf(value)
			objectValueStr := strings.ReplaceAll(objectValue.String(), "\n", "\\\n")
			propertyStrList = append(propertyStrList, prefixWithDOT("")+key+"="+objectValueStr)
		default:
			objectValue := reflect.ValueOf(value).String()
			propertyStrList = append(propertyStrList, prefixWithDOT("")+key+"="+objectValue)
		}
	}
}

func doMapToProperties(propertyStrList []string, value interface{}, prefix string) []string {
	valueKind := reflect.TypeOf(value).Kind()
	switch valueKind {
	case reflect.Map:
		{
			// map结构
			if reflect.ValueOf(value).Len() == 0 {
				return []string{}
			}

			for mapR := reflect.ValueOf(value).MapRange(); mapR.Next(); {
				mapKey := mapR.Key()
				mapValue := mapR.Value()
				propertyStrList = append(propertyStrList, doMapToProperties(propertyStrList, mapValue, prefixWithDOT(prefix)+mapKey.String())...)
			}
		}
	case reflect.Array, reflect.Slice:
		{
			objectValue := reflect.ValueOf(value)
			for index := 0; index < objectValue.Len(); index++ {
				propertyStrList = append(propertyStrList, doMapToProperties(propertyStrList, objectValue.Index(index), prefix+"["+strconv.Itoa(index)+"]")...)
			}
		}
	case reflect.String:
		objectValue := reflect.ValueOf(value)
		objectValueStr := strings.ReplaceAll(objectValue.String(), "\n", "\\\n")
		propertyStrList = append(propertyStrList, prefix+"="+objectValueStr)
	default:
		objectValue := reflect.ValueOf(value).String()
		propertyStrList = append(propertyStrList, prefix+"="+objectValue)
	}
	return propertyStrList
}

func prefixWithDOT(prefix string) string {
	if "" == prefix {
		return ""
	}
	return prefix + "."
}
