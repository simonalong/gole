package util

import (
	"encoding/json"
	"fmt"
	"github.com/simonalong/tools/log"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
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
