package maps

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
	"github.com/simonalong/gole/logger"
	baseTime "github.com/simonalong/gole/time"
	"github.com/simonalong/gole/util"
	"reflect"
	"strings"
	"time"
)

/**
 * 提供新的map
 * 1. 提供类型转换
 * 2. 并发安全
 * 3. 提供有序性
 * 4. 提供与实体的转化功能
 */

// KeyFormat Key格式转换器；示例：
// util.UnderLine：				小驼峰转换为下划线：dataBaseUser -> data_base_user
// util.UnderLineToSmallCamel：	下划线转换为小驼峰：data_base_user -> dataBaseUser
// ...
// 更多函数可以去见cbb-base的util工具中关于字段格式转换的函数
type KeyFormat func(string) string

type GoleMap struct {
	innerMap cmap.ConcurrentMap
	sort     bool
	keys     []string
}

func New() *GoleMap {
	return &GoleMap{
		innerMap: cmap.New(),
		sort:     false,
		keys:     make([]string, 0),
	}
}

func NewSort() *GoleMap {
	return &GoleMap{
		innerMap: cmap.New(),
		sort:     true,
		keys:     make([]string, 0),
	}
}

// Of 支持：k-v-k-v结构
// 默认无序，如果想要有序，请使用OfSort()
func Of(parameters ...any) *GoleMap {
	if parameters == nil || len(parameters) == 0 {
		return New()
	}

	pBaseMap := New()
	for index := 0; index < len(parameters); index++ {
		data := parameters[index]
		if reflect.TypeOf(data).Kind() == reflect.String {
			key := data.(string)
			var value interface{}
			if (index + 1) < len(parameters) {
				value = parameters[index+1]
			}

			pBaseMap.Put(key, value)
			index++
		}
	}
	return pBaseMap
}

func OfSort(parameters ...any) *GoleMap {
	if parameters == nil || len(parameters) == 0 {
		return NewSort()
	}

	pBaseMap := NewSort()
	for index := 0; index < len(parameters); index++ {
		data := parameters[index]
		if reflect.TypeOf(data).Kind() == reflect.String {
			key := data.(string)
			var value interface{}
			if (index + 1) < len(parameters) {
				value = parameters[index+1]
			}

			pBaseMap.Put(key, value)
			index++
		}
	}
	return pBaseMap
}

func From(entity interface{}) (*GoleMap, error) {
	if entity == nil {
		return nil, nil
	}

	// 使用类型断言判断 value 是否为 *GoleMap 类型
	baseMapValue, ok := entity.(*GoleMap)
	if ok {
		return baseMapValue, nil
	}

	entityType := reflect.TypeOf(entity)
	if entityType.Kind() == reflect.Map {
		return FromMap(entity.(map[string]interface{})), nil
	} else if entityType.Kind() == reflect.Struct {
		return FromEntity(entity), nil
	} else if entityType.Kind() == reflect.String {
		return FromJson(entity.(string))
	} else {
		logger.Warnf("暂时不支持除了map、struct和string之外的其他类型：%v", entityType.Kind().String())
		return nil, errors.New(fmt.Sprintf("暂时不支持除了map、struct和string之外的其他类型：%v", entityType.Kind().String()))
	}
}

func FromWithFormat(entity interface{}, keyFormat KeyFormat) (*GoleMap, error) {
	if entity == nil {
		return nil, nil
	}
	// 使用类型断言判断 value 是否为 *GoleMap 类型
	baseMapValue, ok := entity.(*GoleMap)
	if ok {
		return baseMapValue, nil
	}
	entityType := reflect.TypeOf(entity)
	if entityType.Kind() == reflect.Map {
		return FromMapWithFormat(entity.(map[string]interface{}), keyFormat), nil
	} else if entityType.Kind() == reflect.Struct {
		return FromEntityWithFormat(entity, keyFormat), nil
	} else if entityType.Kind() == reflect.String {
		return FromJsonWithFormat(entity.(string), keyFormat)
	} else {
		logger.Warnf("暂时不支持除了map、struct和string之外的其他类型：%v", entityType.Kind().String())
		return nil, errors.New(fmt.Sprintf("暂时不支持除了map、struct和string之外的其他类型：%v", entityType.Kind().String()))
	}
}

func FromRows(rows driver.Rows) []*GoleMap {
	if rows == nil {
		return []*GoleMap{}
	}
	columns := rows.Columns()
	dest := make([]driver.Value, len(columns))
	var ormMapList []*GoleMap
	for rows.Next(dest) == nil {
		ormMap := NewSort()
		for index, column := range columns {
			ormMap.Put(column, dest[index])
		}
		ormMapList = append(ormMapList, ormMap)
	}
	return ormMapList
}

func FromSqlRows(rows *sql.Rows) []*GoleMap {
	if rows == nil {
		return []*GoleMap{}
	}

	columns, err := rows.Columns()
	if err != nil {
		logger.Errorf("获取columns异常：%v", err)
		return []*GoleMap{}
	}

	var ormMapList []*GoleMap
	for rows.Next() {
		ptrs := make([]interface{}, len(columns))
		container := make([]interface{}, len(columns))
		for i := range ptrs {
			ptrs[i] = &container[i]
		}
		if err := rows.Scan(ptrs...); err != nil {
			logger.Errorf("scan转换字段异常：%v", err)
			return []*GoleMap{}
		}

		ormMap := OfSort()
		for index, columnName := range columns {
			if container[index] == nil {
				continue
			}
			ormMap.Put(columnName, container[index])
		}
		ormMapList = append(ormMapList, ormMap)
	}
	return ormMapList
}

func FromRowsWithFormat(rows driver.Rows, keyFormat KeyFormat) []*GoleMap {
	if rows == nil {
		return []*GoleMap{}
	}
	columns := rows.Columns()
	dest := make([]driver.Value, len(columns))
	var ormMapList []*GoleMap
	for rows.Next(dest) == nil {
		ormMap := NewSort()
		for index, column := range columns {
			ormMap.Put(keyFormat(column), dest[index])
		}
		ormMapList = append(ormMapList, ormMap)
	}
	return ormMapList
}

func FromEntityList(entityList []interface{}) []*GoleMap {
	if len(entityList) == 0 {
		return []*GoleMap{}
	}

	var ormMapList []*GoleMap
	for _, entity := range entityList {
		ormMapList = append(ormMapList, FromEntity(entity))
	}
	return ormMapList
}

// FromEntity 从实体转换为map，默认转换为有序map
func FromEntity(entity interface{}) *GoleMap {
	if entity == nil {
		return New()
	}

	// 使用类型断言判断 value 是否为 *GoleMap 类型
	baseMapValue, ok := entity.(*GoleMap)
	if ok {
		return baseMapValue
	}

	objType := reflect.TypeOf(entity)
	// 只接收结构体类型
	if objType.Kind() != reflect.Struct {
		return nil
	}

	entityMap := NewSort()

	objValue := reflect.ValueOf(entity)
	for fieldIndex, num := 0, objType.NumField(); fieldIndex < num; fieldIndex++ {
		field := objType.Field(fieldIndex)
		if !util.IsPublic(field.Name) {
			continue
		}

		columnName := getFinalColumnName(field)

		fieldValue := objValue.Field(fieldIndex)
		entityMap.Put(columnName, fieldValue.Interface())
	}
	return entityMap
}

// FromEntityWithFormat 从实体转换为map，默认转换为有序map
func FromEntityWithFormat(entity interface{}, keyFormat KeyFormat) *GoleMap {
	if entity == nil {
		return New()
	}

	objType := reflect.TypeOf(entity)
	// 只接收结构体类型
	if objType.Kind() != reflect.Struct {
		return nil
	}

	entityMap := NewSort()

	objValue := reflect.ValueOf(entity)
	for fieldIndex, num := 0, objType.NumField(); fieldIndex < num; fieldIndex++ {
		field := objType.Field(fieldIndex)
		if !util.IsPublic(field.Name) {
			continue
		}

		columnName := getFinalColumnName(field)

		fieldValue := objValue.Field(fieldIndex)
		entityMap.Put(keyFormat(columnName), fieldValue.Interface())
	}
	return entityMap
}

// FromMap 从map转换为OrmMap，默认转换为有序map
func FromMap(dataMap map[string]interface{}) *GoleMap {
	if dataMap == nil || len(dataMap) == 0 {
		return New()
	}
	resultMap := NewSort()
	for key, val := range dataMap {
		resultMap.Put(key, val)
	}
	return resultMap
}

func FromMapWithFormat(dataMap map[string]interface{}, keyFormat KeyFormat) *GoleMap {
	if dataMap == nil || len(dataMap) == 0 {
		return New()
	}
	resultMap := NewSort()
	for key, val := range dataMap {
		resultMap.Put(keyFormat(key), val)
	}
	return resultMap
}

func FromJson(jsonOfContent string) (*GoleMap, error) {
	if jsonOfContent == "" {
		return New(), nil
	}
	resultMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonOfContent), &resultMap)
	if err != nil {
		logger.Warnf("JsonToMap, error: %v, content: %v", err, jsonOfContent)
		return nil, err
	}

	return FromMap(resultMap), nil
}

func FromJsonWithFormat(jsonOfContent string, keyFormat KeyFormat) (*GoleMap, error) {
	if jsonOfContent == "" {
		return New(), nil
	}
	resultMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonOfContent), &resultMap)
	if err != nil {
		logger.Warnf("JsonToMap, error: %v, content: %v", err, jsonOfContent)
		return nil, err
	}

	return FromMapWithFormat(resultMap, keyFormat), nil
}

// AllIsEmpty 所有数据都为空，则返回true
func AllIsEmpty(dataMaps []*GoleMap) bool {
	if dataMaps == nil {
		return true
	}

	for _, dataMap := range dataMaps {
		if dataMap.IsUnEmpty() {
			return false
		}
	}
	return true
}

func (receiver *GoleMap) AsDeepMap() *GoleMap {
	dataMap1 := receiver.ToMap()
	properties, err := util.MapToProperties(dataMap1)
	if err != nil {
		logger.Error("AsDeepMap，MapToProperties, error: %v", err)
		return nil
	}

	dataMap, err := util.PropertiesToMap(properties)
	if err != nil {
		logger.Error("AsDeepMap，PropertiesToMap, error: %v", err)
		return nil
	}
	return FromMap(dataMap)
}

func (receiver *GoleMap) ToEntity(pEntity interface{}) error {
	if pEntity == nil {
		return errors.New("对象指针为nil")
	}
	targetType := reflect.TypeOf(pEntity)
	if targetType.Kind() != reflect.Ptr {
		return errors.New("暂时只支持：待转换对向要为指针类型")
	}

	if targetType.Elem().Kind() != reflect.Struct {
		return errors.New("暂时只支持：待转换对向指针指向的类型要为实体类型")
	}

	targetValue := reflect.ValueOf(pEntity)
	for index, num := 0, targetType.Elem().NumField(); index < num; index++ {
		field := targetType.Elem().Field(index)
		fieldValue := targetValue.Elem().Field(index)

		invokeValue(receiver.innerMap.Items(), field, fieldValue)
	}

	return nil
}

func (receiver *GoleMap) ToMap() map[string]interface{} {
	dataMap := receiver.innerMap.Items()
	rsultMap := map[string]interface{}{}
	for k, v := range dataMap {
		if reflect.ValueOf(v).Kind() != reflect.Ptr {
			rsultMap[k] = v
			continue
		}
		if mapV := v.(*GoleMap); mapV != nil {
			rsultMap[k] = mapV.ToMap()
		} else {
			rsultMap[k] = v
		}
	}
	return rsultMap
}

func (receiver *GoleMap) CloneExceptKeys(keys []string) *GoleMap {
	resultMap := receiver.Clone()
	resultMap.RemoveKeys(keys)
	return resultMap
}

func (receiver *GoleMap) ToJson() string {
	return util.ToJsonString(receiver.innerMap)
}

func (receiver *GoleMap) ToJsonOfSort() string {
	if receiver == nil || receiver.IsEmpty() {
		return "{}"
	}
	var jsonResult string
	jsonResult += "{"

	var kvs []string
	for _, key := range receiver.Keys() {
		val, _ := receiver.Get(key)
		if val == nil {
			continue
		}
		valType := reflect.TypeOf(val)
		valValue := reflect.ValueOf(val)
		if valType.Kind() == reflect.Ptr {
			valType = valType.Elem()
			valValue = valValue.Elem()
		}
		if util.IsStringType(valType) {
			kvs = append(kvs, fmt.Sprintf("\"%v\":\"%v\"", key, valValue.Interface()))
		} else if util.IsNumberType(valType) || util.IsBoolType(valType) {
			kvs = append(kvs, fmt.Sprintf("\"%v\":%v", key, valValue.Interface()))
		} else if util.IsTimeType(valType) {
			kvs = append(kvs, fmt.Sprintf("\"%v\":\"%v\"", key, valValue.Interface()))
		} else {
			var valJson string
			if valType == reflect.TypeOf(GoleMap{}) {
				value := val.(*GoleMap)
				valJson = value.ToJsonOfSort()
			} else {
				valJson = FromEntity(valValue.Interface()).ToJsonOfSort()
			}
			if valJson == "{}" {
				continue
			}
			kvs = append(kvs, fmt.Sprintf("\"%v\":%v", key, valJson))
		}
	}

	jsonResult += strings.Join(kvs, ",")
	jsonResult += "}"
	return jsonResult
}

func (receiver *GoleMap) ToString() string {
	var keyValue []string
	for _, key := range receiver.Keys() {
		val, _ := receiver.Get(key)
		if val == nil {
			continue
		}
		valType := reflect.TypeOf(val)
		if util.IsStringType(valType) {
			keyValue = append(keyValue, "\""+key+"\":\""+util.ToString(val)+"\"")
		} else {
			keyValue = append(keyValue, "\""+key+"\":"+util.ToString(val)+"")
		}
	}
	return "[" + strings.Join(keyValue, ",") + "]"
}

func (receiver *GoleMap) Keys() []string {
	if receiver.sort {
		return receiver.keys
	} else {
		return receiver.innerMap.Keys()
	}
}

func (receiver *GoleMap) Values() []interface{} {
	if receiver.sort {
		keys := receiver.keys
		var valueList []interface{}
		for _, key := range keys {
			val, _ := receiver.Get(key)
			valueList = append(valueList, val)
		}
		return valueList
	} else {
		keys := receiver.innerMap.Keys()
		var valueList []interface{}
		for _, key := range keys {
			val, _ := receiver.Get(key)
			valueList = append(valueList, val)
		}
		return valueList
	}
}

// SetSort 设置map为有序或者无序map
// 注意：
//  1. 如果从无序变为有序，且之前已经有一些数据，则之前的数据顺序至此固定，后续的顺序就按照添加的顺序固定
//  2. 如果从有序变为无序，且之前已经有一些数据，则顺序就完全乱掉了
func (receiver *GoleMap) SetSort(sort bool) *GoleMap {
	if !receiver.sort && sort {
		receiver.keys = receiver.innerMap.Keys()
	} else if receiver.sort && !sort {
		receiver.keys = make([]string, 0)
	}
	receiver.sort = sort
	return receiver
}

func (receiver *GoleMap) IsEmpty() bool {
	return len(receiver.innerMap.Keys()) == 0
}

func (receiver *GoleMap) IsUnEmpty() bool {
	return len(receiver.innerMap.Keys()) != 0
}

func (receiver *GoleMap) Clone() *GoleMap {
	cloneMap := &GoleMap{
		innerMap: cmap.New(),
		sort:     receiver.sort,
		keys:     make([]string, 0),
	}
	for _, key := range receiver.Keys() {
		val, _ := receiver.Get(key)
		cloneMap.Put(key, val)
	}
	return cloneMap
}

func (receiver *GoleMap) Put(key string, value interface{}) *GoleMap {
	if key == "" {
		return receiver
	}
	receiver.innerMap.Set(key, value)
	if receiver.sort {
		if !util.ListContains(receiver.keys, key) {
			receiver.keys = append(receiver.keys, key)
		}
	}
	return receiver
}

func (receiver *GoleMap) Set(key string, value interface{}) *GoleMap {
	return receiver.Put(key, value)
}

func (receiver *GoleMap) Contain(key string) bool {
	if key == "" {
		return false
	}
	_, exit := receiver.innerMap.Get(key)
	return exit
}

func (receiver *GoleMap) Get(key string) (interface{}, bool) {
	if key == "" {
		return nil, false
	}
	return receiver.innerMap.Get(key)
}

func (receiver *GoleMap) GetInt(key string) (int, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToInt(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetInt8(key string) (int8, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToInt8(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetInt16(key string) (int16, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToInt16(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetInt32(key string) (int32, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToInt32(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetInt64(key string) (int64, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToInt64(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetUInt(key string) (uint, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToUInt(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetUInt8(key string) (uint8, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToUInt8(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetUInt16(key string) (uint16, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToUInt16(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetUInt32(key string) (uint32, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToUInt32(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetUInt64(key string) (uint64, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToUInt64(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetFloat32(key string) (float32, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToFloat32(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetFloat64(key string) (float64, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToFloat64(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetBool(key string) (bool, bool) {
	if key == "" {
		return false, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToBool(d), true
	} else {
		return false, false
	}
}

func (receiver *GoleMap) GetComplex64(key string) (complex64, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToComplex64(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetComplex128(key string) (complex128, bool) {
	if key == "" {
		return 0, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return util.ToComplex128(d), true
	} else {
		return 0, false
	}
}

func (receiver *GoleMap) GetString(key string) (string, bool) {
	if key == "" {
		return "", false
	}
	val, exit := receiver.innerMap.Get(key)
	if exit {
		if timestampVal, ok := val.(time.Time); ok {
			return baseTime.TimeToStringYmdHmsS(timestampVal), true
		} else {
			return util.ToString(val), true
		}
	} else {
		return "", false
	}
}

func (receiver *GoleMap) GetTime(key string) (time.Time, bool) {
	if key == "" {
		return time.Time{}, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		if d == nil {
			return time.Time{}, false
		}
		if reflect.TypeOf(d).Kind() == reflect.String {
			dataTime, err := baseTime.ParseTime(d.(string))
			if err != nil {
				return time.Time{}, false
			}
			return dataTime, true
		}
		return d.(time.Time), true
	} else {
		logger.Warnf("map中的key（%v）获取time不存在", key)
		return time.Now(), false
	}
}

func (receiver *GoleMap) GetBytes(key string) ([]byte, bool) {
	if key == "" {
		return make([]byte, 0), false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		return []byte(util.ToString(d)), true
	} else {
		return nil, false
	}
}

func (receiver *GoleMap) GetMaps(key string) (*GoleMap, bool) {
	if key == "" {
		return nil, false
	}
	d, exit := receiver.innerMap.Get(key)
	if exit {
		val, err := From(d)
		if err != nil {
			return nil, false
		}
		return val, true
	} else {
		return nil, false
	}
}

func (receiver *GoleMap) Remove(key string) {
	receiver.innerMap.Remove(key)
	if receiver.sort {
		id := util.IndexOf(receiver.keys, key)
		receiver.keys = append(receiver.keys[:id], receiver.keys[id+1:]...)
	}
}

func (receiver *GoleMap) RemoveKeys(keys []string) {
	for _, key := range keys {
		receiver.innerMap.Remove(key)
	}

	if receiver.sort {
		for _, key := range keys {
			id := util.IndexOf(receiver.keys, key)
			receiver.keys = append(receiver.keys[:id], receiver.keys[id+1:]...)
		}
	}
}

func (receiver *GoleMap) RemoveAll() {
	receiver.innerMap.Clear()
	if receiver.sort {
		receiver.keys = make([]string, 0)
	}
}
func (receiver *GoleMap) Clear() {
	receiver.innerMap.Clear()
	if receiver.sort {
		receiver.keys = make([]string, 0)
	}
}
func (receiver *GoleMap) Size() int {
	return len(receiver.innerMap.Keys())
}

func getFinalColumnName(field reflect.StructField) string {
	// 先读取标签column
	columnName := field.Tag.Get("column")
	if len(columnName) != 0 {
		return columnName
	}

	// 如果没有配置column标签，也可以使用json标签，这里也支持
	aliasJson := field.Tag.Get("json")
	if len(aliasJson) != 0 {
		index := strings.Index(aliasJson, ",")
		if index != -1 {
			return aliasJson[:index]
		} else {
			return aliasJson
		}
	}

	// 如果也没有配置json标签，则使用属性的属性名，将首字母变小写
	return util.ToLowerFirstPrefix(field.Name)
}

func invokeValue(srcMap map[string]interface{}, field reflect.StructField, fieldValue reflect.Value) {
	if srcMap == nil {
		return
	}
	// 私有字段不处理
	if util.IsPrivate(field.Name) {
		return
	}

	var srcValue interface{}
	// 优先使用 column
	aliasJson := field.Tag.Get("column")
	index := strings.Index(aliasJson, ",")
	if index != -1 {
		aliasJson = aliasJson[:index]
	}
	if v, exist := srcMap[aliasJson]; exist {
		// 使用标签：column
		srcValue = v
	} else {
		aliasJson := field.Tag.Get("json")
		index := strings.Index(aliasJson, ",")
		if index != -1 {
			aliasJson = aliasJson[:index]
		}
		if v, exist = srcMap[aliasJson]; exist {
			// 兼容标签：json
			srcValue = v
		} else if v, exist := srcMap[util.BigCamelToSmallCamel(field.Name)]; exist {
			// 兼容dataSeatakUser格式读取
			srcValue = v
		} else {
			// 其他格式暂时都不支持
			return
		}
	}

	srcVal := reflect.ValueOf(srcValue)
	targetValue := util.ValueToTarget(srcVal, field.Type)
	if targetValue.IsValid() {
		fieldValue.Set(targetValue.Convert(field.Type))
	}
}
