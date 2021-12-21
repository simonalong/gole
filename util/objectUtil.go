package util

import (
	"encoding/json"
	"fmt"
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

var utilLog *logrus.Logger

func init() {
	utilLog = log.GetLogger("utilLog")
}

func ObjectToMap(data interface{}) map[string]interface{} {
	if reflect.TypeOf(data).Kind() == reflect.Map {
		resultMap := map[string]interface{}{}
		dataValue := reflect.ValueOf(data)
		for mapR := dataValue.MapRange(); mapR.Next(); {
			mapKey := mapR.Key()
			mapValue := mapR.Value()

			resultMap[mapKey.String()] = mapValue
		}
		return resultMap
	}
	return nil
}

func IsNumber(fieldKing reflect.Kind) bool {
	switch fieldKing {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	default:
		return false
	}
}

// IsBaseType 是否是常见基本类型
func IsBaseType(fieldType reflect.Type) bool {
	fieldKind := fieldType.Kind()
	if fieldKind == reflect.Ptr {
		fieldKind = fieldType.Elem().Kind()
	}

	switch fieldKind {
	case reflect.Int:
		return true
	case reflect.Int8:
		return true
	case reflect.Int16:
		return true
	case reflect.Int32:
		return true
	case reflect.Int64:
		return true
	case reflect.Uint:
		return true
	case reflect.Uint8:
		return true
	case reflect.Uint16:
		return true
	case reflect.Uint32:
		return true
	case reflect.Uint64:
		return true
	case reflect.Float32:
		return true
	case reflect.Float64:
		return true
	case reflect.Bool:
		return true
	case reflect.String:
		return true
	default:
		if fieldType.String() == "time.Time" {
			return true
		}
		return false
	}
}

func ToJsonString(data interface{}) string {
	bytes, _ := json.Marshal(data)
	return string(bytes)
}

func ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

func ToInt(value interface{}) int {
	result, err := ToValue(value, reflect.Int)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(int)
}

func ToInt8(value interface{}) int8 {
	result, err := ToValue(value, reflect.Int8)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(int8)
}

func ToInt16(value interface{}) int16 {
	result, err := ToValue(value, reflect.Int16)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(int16)
}

func ToInt32(value interface{}) int32 {
	result, err := ToValue(value, reflect.Int32)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(int32)
}

func ToInt64(value interface{}) int64 {
	result, err := ToValue(value, reflect.Int64)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(int64)
}

func ToUInt(value interface{}) uint {
	result, err := ToValue(value, reflect.Uint)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(uint)
}

func ToUInt8(value interface{}) uint8 {
	result, err := ToValue(value, reflect.Uint8)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(uint8)
}

func ToUInt16(value interface{}) uint16 {
	result, err := ToValue(value, reflect.Uint16)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(uint16)
}

func ToUInt32(value interface{}) uint32 {
	result, err := ToValue(value, reflect.Uint32)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(uint32)
}

func ToUInt64(value interface{}) uint64 {
	result, err := ToValue(value, reflect.Uint64)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(uint64)
}

func ToFloat32(value interface{}) float32 {
	result, err := ToValue(value, reflect.Float32)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(float32)
}

func ToFloat64(value interface{}) float64 {
	result, err := ToValue(value, reflect.Float64)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return 0
	}
	return result.(float64)
}

func ToBool(value interface{}) bool {
	result, err := ToValue(value, reflect.Bool)
	if err != nil {
		utilLog.Errorf("%v", err.Error())
		return false
	}
	return result.(bool)
}

func ToValue(value interface{}, valueKind reflect.Kind) (interface{}, error) {
	valueStr := ToString(value)
	return Cast(valueKind, valueStr)
}

func Cast(fieldKind reflect.Kind, valueStr string) (interface{}, error) {
	if valueStr == "nil" || valueStr == "" {
		return nil, nil
	}
	switch fieldKind {
	case reflect.Int:
		return strconv.Atoi(valueStr)
	case reflect.Int8:
		return strconv.ParseInt(valueStr, 10, 8)
	case reflect.Int16:
		return strconv.ParseInt(valueStr, 10, 16)
	case reflect.Int32:
		return strconv.ParseInt(valueStr, 10, 32)
	case reflect.Int64:
		return strconv.ParseInt(valueStr, 10, 64)
	case reflect.Uint:
		return strconv.ParseUint(valueStr, 10, 0)
	case reflect.Uint8:
		return strconv.ParseUint(valueStr, 10, 8)
	case reflect.Uint16:
		return strconv.ParseUint(valueStr, 10, 16)
	case reflect.Uint32:
		return strconv.ParseUint(valueStr, 10, 32)
	case reflect.Uint64:
		return strconv.ParseUint(valueStr, 10, 64)
	case reflect.Float32:
		return strconv.ParseFloat(valueStr, 32)
	case reflect.Float64:
		return strconv.ParseFloat(valueStr, 64)
	case reflect.Bool:
		return strconv.ParseBool(valueStr)
	}
	return valueStr, nil
}

func MapToObject(dataMap map[string]interface{}, targetObj interface{}) {
	targetType := reflect.TypeOf(targetObj)
	if targetType.Kind() != reflect.Ptr {
		utilLog.Warn("targetObj type is not ptr")
		return
	}

	targetValue := reflect.ValueOf(targetObj)
	for index, num := 0, targetType.Elem().NumField(); index < num; index++ {
		field := targetType.Elem().Field(index)
		fieldValue := targetValue.Elem().Field(index)

		doInvoke(dataMap, field, fieldValue)
	}
}

func doInvoke(dataMap map[string]interface{}, structField reflect.StructField, fieldValue reflect.Value) {
	// 私有字段不处理
	if !isStartUpper(structField.Name) {
		return
	}

	fieldKind := structField.Type.Kind()

	// 基本类型
	if IsBaseType(structField.Type) {
		if v, exist := getValue(dataMap, structField.Name); exist {
			if fieldValue.Kind() == reflect.Ptr {
				fieldValue.Elem().FieldByName(structField.Name).Set(reflect.ValueOf(v))
			} else {
				fieldValue.Set(reflect.ValueOf(v))
			}
		}
	} else if fieldKind == reflect.Struct {
		// 结构体类型
		if v, exist := getValue(dataMap, structField.Name); exist {
			if reflect.TypeOf(v).Kind() == reflect.Map {
				fieldValueTem := reflect.New(structField.Type)
				fMapValue := reflect.ValueOf(v)

				for mapR := fMapValue.MapRange(); mapR.Next(); {
					mapKey := mapR.Key()
					mapValue := mapR.Value()

					doInvoke(fieldDataMap, structField.Type.Field(fIndex), fieldValueTem)
					fieldValueMap.SetMapIndex(mapKey, mapValue)
				}

				for fIndex, fNum := 0, structField.Type.NumField(); fIndex < fNum; fIndex++ {
					doInvoke(fieldDataMap, structField.Type.Field(fIndex), fieldValueTem)
				}
				d := fieldValueTem.Elem()

				if fieldValue.Kind() == reflect.Ptr {
					fieldValue.Elem().FieldByName(structField.Name).Set(d)
				} else {
					fieldValue.Set(d)
				}
			}
		}
	} else if fieldKind == reflect.Map {
		// map结构
		if v, exist := getValue(dataMap, structField.Name); exist {
			if reflect.TypeOf(v).Kind() == reflect.Map {
				fieldValueMap := reflect.MakeMap(structField.Type)
				fMapValue := reflect.ValueOf(v)

				// map结构
				if fMapValue.Len() == 0 {
					return
				}

				for mapR := fMapValue.MapRange(); mapR.Next(); {
					mapKey := mapR.Key()
					mapValue := mapR.Value()

					doInvoke(fieldDataMap, structField.Type.Field(fIndex), fieldValueTem)
					fieldValueMap.SetMapIndex(mapKey, mapValue)
				}

				if fieldValue.Kind() == reflect.Ptr {
					fieldValue.Elem().FieldByName(structField.Name).Set(fieldValueMap)
				} else {
					fieldValue.Set(fieldValueMap)
				}
			}
		}
	} else if fieldKind == reflect.Array || fieldKind == reflect.Slice {
		// 数组结构

	} else {

	}
}

func getValue(dataMap map[string]interface{}, key string) (interface{}, bool) {
	if v1, exits := dataMap[key]; exits {
		return v1, true
	} else if v2, exits := dataMap[ToLowerFirstPrefix(key)]; exits {
		return v2, true
	}
	return nil, false
}

func doMapToObject(dataMap map[string]interface{}, numField int, targetObj interface{}) {
	targetType := reflect.TypeOf(targetObj)
	targetValue := reflect.ValueOf(targetObj)

	if targetType.Kind() != reflect.Ptr {
		utilLog.Warn("targetObj type is not ptr")
		return
	}

	targetTypeE := targetType.Elem()
	if targetTypeE.Kind() == reflect.Ptr {
		targetTypeE = targetType.Elem()
	}

	if targetTypeE.Kind() == reflect.Struct {
		fmt.Println("hahaha")
	}

	for index, num := 0, numField; index < num; index++ {
		field := targetTypeE.Field(index)
		fieldValue := targetValue.Elem().Field(index)

		// 私有字段不处理
		if !isStartUpper(field.Name) {
			continue
		}

		fieldKind := field.Type.Kind()

		// 基本类型
		if IsBaseType(field.Type) {
			if v1, exits := dataMap[field.Name]; exits {
				fieldValue.Set(reflect.ValueOf(v1))
			} else if v2, exits := dataMap[ToLowerFirstPrefix(field.Name)]; exits {
				fieldValue.Set(reflect.ValueOf(v2))
			}
		} else if fieldKind == reflect.Struct {
			// 结构体类型
			if v1, exits := dataMap[field.Name]; exits {
				if reflect.TypeOf(v1).Kind() == reflect.Map {
					fieldValueTem := reflect.New(field.Type)
					fieldDataValue := fieldValueTem.Interface()
					MapToObject(v1.(map[string]interface{}), &fieldDataValue)
					fieldValue.Set(reflect.ValueOf(fieldDataValue))
				}
			} else if v2, exits := dataMap[ToLowerFirstPrefix(field.Name)]; exits {
				if reflect.TypeOf(v2).Kind() == reflect.Map {
					fieldValueTem := reflect.New(field.Type)
					for index, num := 0, field.Type.NumField(); index < num; index++ {
						fmt.Println("asdf")
					}

					fieldDataValue := fieldValueTem.Interface()
					MapToObject(v2.(map[string]interface{}), &fieldDataValue)
					fieldValue.Set(reflect.ValueOf(fieldDataValue))
				}
			}
		} else if fieldKind == reflect.Map {
			// map结构

		} else if fieldKind == reflect.Array || fieldKind == reflect.Slice {
			// 数组结构
		} else {

		}
	}
}

// 判断首字母是否大写
func isStartUpper(s string) bool {
	return unicode.IsUpper([]rune(s)[0])
}

// ToLowerFirstPrefix 首字母小写
func ToLowerFirstPrefix(dataStr string) string {
	return strings.ToLower(dataStr[:1]) + dataStr[1:]
}

// ToUpperFirstPrefix 首字母大写
func ToUpperFirstPrefix(dataStr string) string {
	return strings.ToLower(dataStr[:1]) + dataStr[1:]
}
