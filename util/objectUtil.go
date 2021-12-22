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

		doInvokeValue(reflect.ValueOf(dataMap), field, fieldValue)
	}
}

func doInvokeValue(fieldMapValue reflect.Value, field reflect.StructField, fieldValue reflect.Value) {
	// 私有字段不处理
	if !isStartUpper(field.Name) {
		return
	}

	if fieldMapValue.Kind() == reflect.Ptr {
		fieldMapValue = fieldMapValue.Elem()
	}

	if v, exist := getValueFromMapValue(fieldMapValue, field.Name); exist {
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		targetValue := valueToTarget(v, field.Type)
		if targetValue.IsValid() {
			if fieldValue.Kind() == reflect.Ptr {
				if targetValue.Kind() == reflect.Ptr {
					fieldValue.Elem().FieldByName(field.Name).Set(targetValue.Elem())
				} else {
					fieldValue.Elem().FieldByName(field.Name).Set(targetValue)
				}
			} else {
				if targetValue.Kind() == reflect.Ptr {
					fieldValue.Set(targetValue.Elem())
				} else {
					fieldValue.Set(targetValue)
				}
			}
		}
	}
}

func valueToTarget(srcValue reflect.Value, dstType reflect.Type) reflect.Value {
	if dstType.Kind() == reflect.Struct {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if sourceValue.Kind() == reflect.Map || sourceValue.Kind() == reflect.Struct {
			mapFieldValue := reflect.New(dstType)
			for index, num := 0, mapFieldValue.Type().Elem().NumField(); index < num; index++ {
				field := mapFieldValue.Type().Elem().Field(index)
				fieldValue := mapFieldValue.Elem().Field(index)

				doInvokeValue(sourceValue, field, fieldValue)
			}
			return mapFieldValue
		}
	} else if dstType.Kind() == reflect.Map {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if sourceValue.Kind() == reflect.Map {
			mapFieldValue := reflect.MakeMap(dstType)
			for mapR := sourceValue.MapRange(); mapR.Next(); {
				mapKey := mapR.Key()
				mapValue := mapR.Value()

				mapKeyRealValue, err := Cast(mapFieldValue.Type().Key().Kind(), fmt.Sprintf("%v", mapKey.Interface()))
				mapValueRealValue := valueToTarget(mapValue, mapFieldValue.Type().Elem())
				if err == nil {
					if mapValueRealValue.Kind() == reflect.Ptr {
						mapFieldValue.SetMapIndex(reflect.ValueOf(mapKeyRealValue), mapValueRealValue.Elem())
					} else {
						mapFieldValue.SetMapIndex(reflect.ValueOf(mapKeyRealValue), mapValueRealValue)
					}
				}
			}
			return mapFieldValue
		}
	} else if dstType.Kind() == reflect.Slice || dstType.Kind() == reflect.Array {
		if srcValue.Kind() == reflect.Ptr {
			srcValue = srcValue.Elem()
		}
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if sourceValue.Kind() == reflect.Slice || sourceValue.Kind() == reflect.Array {
			arrayFieldValue := reflect.MakeSlice(dstType, 0, 0)
			for arrayIndex := 0; arrayIndex < sourceValue.Len(); arrayIndex++ {
				dataV := valueToTarget(sourceValue.Index(arrayIndex), dstType.Elem())
				if dataV.IsValid() {
					if dataV.Kind() == reflect.Ptr {
						arrayFieldValue = reflect.Append(arrayFieldValue, dataV.Elem())
					} else {
						arrayFieldValue = reflect.Append(arrayFieldValue, dataV)
					}
				}
			}
			return arrayFieldValue
		}
	} else if IsBaseType(dstType) {
		sourceValue := reflect.ValueOf(srcValue.Interface())
		if IsBaseType(sourceValue.Type()) {
			v, err := Cast(dstType.Kind(), fmt.Sprintf("%v", srcValue.Interface()))
			if err == nil {
				return reflect.ValueOf(v)
			}
		}
	} else {
		return reflect.ValueOf(nil)
	}
	return reflect.ValueOf(nil)
}

func getValueFromMap(dataMap map[string]interface{}, key string) (interface{}, bool) {
	//val := dataMap.MapIndex(reflect.ValueOf(key))
	//if val != nil {
	//
	//}
	if v1, exits := dataMap[key]; exits {
		return v1, true
	} else if v2, exits := dataMap[ToLowerFirstPrefix(key)]; exits {
		return v2, true
	}
	return nil, false
}

func getValueFromMapValue(keyValues reflect.Value, key string) (reflect.Value, bool) {
	if keyValues.Kind() == reflect.Map {
		if v1 := keyValues.MapIndex(reflect.ValueOf(key)); v1.IsValid() {
			return v1, true
		} else if v2 := keyValues.MapIndex(reflect.ValueOf(ToLowerFirstPrefix(key))); v2.IsValid() {
			return v2, true
		}
	} else if keyValues.Kind() == reflect.Struct {
		if v1 := keyValues.FieldByName(key); v1.IsValid() {
			return v1, true
		} else if v2 := keyValues.FieldByName(ToLowerFirstPrefix(key)); v2.IsValid() {
			return v2, true
		}
	}

	return reflect.ValueOf(nil), false
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
