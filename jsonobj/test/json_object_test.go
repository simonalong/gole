package test

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/simonalong/gole/jsonobj"
	"github.com/simonalong/gole/util"
	"testing"
)

//		{
//	   "store": {
//	       "book": [
//	           {
//	               "category": "reference",
//	               "author": "Nigel Rees",
//	               "title": "Sayings of the Century",
//	               "price": 8.95
//	           },
//	           {
//	               "category": "fiction",
//	               "author": "Evelyn Waugh",
//	               "title": "Sword of Honour",
//	               "price": 12.99
//	           },
//	           {
//	               "category": "fiction",
//	               "author": "Herman Melville",
//	               "title": "Moby Dick",
//	               "isbn": "0-553-21311-3",
//	               "price": 8.99
//	           },
//	           {
//	               "category": "fiction",
//	               "author": "J. R. R. Tolkien",
//	               "title": "The Lord of the Rings",
//	               "isbn": "0-395-19395-8",
//	               "price": 22.99
//	           }
//	       ],
//	       "bicycle": {
//	           "color": "red",
//	           "price": 19.95
//	       }
//	   },
//	   "expensive": 10
//	}
var str = "{\n    \"store\": {\n        \"book\": [\n            {\n                \"category\": \"reference\",\n                \"author\": \"Nigel Rees\",\n                \"title\": \"Sayings of the Century\",\n                \"price\": 8.95\n            },\n            {\n                \"category\": \"fiction\",\n                \"author\": \"Evelyn Waugh\",\n                \"title\": \"Sword of Honour\",\n                \"price\": 12.99\n            },\n            {\n                \"category\": \"fiction\",\n                \"author\": \"Herman Melville\",\n                \"title\": \"Moby Dick\",\n                \"isbn\": \"0-553-21311-3\",\n                \"price\": 8.99\n            },\n            {\n                \"category\": \"fiction\",\n                \"author\": \"J. R. R. Tolkien\",\n                \"title\": \"The Lord of the Rings\",\n                \"isbn\": \"0-395-19395-8\",\n                \"price\": 22.99\n            }\n        ],\n        \"bicycle\": {\n            \"color\": \"red\",\n            \"price\": 19.95\n        }\n    },\n    \"expensive\": 10\n}"

func TestLoad(t *testing.T) {
	// 将JSON数据解析为interface{}
	jsonObj, err := jsonobj.Load(str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	assert.Equal(t, jsonObj.GetString("$.expensive"), util.ToString(10))
	assert.Equal(t, jsonObj.GetString("$.store.book[0].price"), util.ToString(8.95))
	assert.Equal(t, jsonObj.GetString("$.store.book[-1].isbn"), util.ToString("0-395-19395-8"))
	assert.Equal(t, jsonObj.GetString("$.store.book[0,1].price"), util.ToString("[8.95 12.99]"))
	assert.Equal(t, jsonObj.GetString("$.store.book[0:2].price"), util.ToString("[8.95 12.99 8.99]"))
	assert.Equal(t, jsonObj.GetString("$.store.book[?(@.isbn)].price"), util.ToString("[8.99 22.99]"))
	assert.Equal(t, jsonObj.GetString("$.store.book[?(@.isbn > 10)].title"), util.ToString("[Sayings of the Century Sword of Honour]"))
}

//	{
//	   "intValue": 12,
//	   "intValue8": 12,
//	   "intValue16": 12,
//	   "intValue32": 12,
//	   "intValue64": 12,
//	   "stringValue": "haode",
//	   "boolValue": false,
//	   "objectValue": {
//	       "field1": true,
//	       "field2Struct": {
//	           "f21": 43,
//	           "k2array": [
//	               1,
//	               2,
//	               3,
//	               4
//	           ]
//	       }
//	   }
//	}
func TestPut4(t *testing.T) {
	str := "{\n    \"intValue\":12,\n    \"intValue8\":12,\n    \"intValue16\":12,\n    \"intValue32\":12,\n    \"intValue64\":12,\n    \"stringValue\":\"haode\",\n    \"boolValue\":false,\n    \"objectValue\":{\n        \"field1\":true,\n        \"field2Struct\":{\n            \"f21\":43,\n            \"k2array\":[\n                1,\n                2,\n                3,\n                4\n            ]\n        }\n    }\n}"
	jsonObj, err := jsonobj.Load(str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	assert.Equal(t, jsonObj.GetInt("$.intValue"), 12)
	assert.Equal(t, jsonObj.GetInt8("$.intValue"), int8(12))
	assert.Equal(t, jsonObj.GetInt16("$.intValue"), int16(12))
	assert.Equal(t, jsonObj.GetInt32("$.intValue"), int32(12))
	assert.Equal(t, jsonObj.GetInt64("$.intValue"), int64(12))

	assert.Equal(t, jsonObj.GetString("$.stringValue"), "haode")
	assert.Equal(t, jsonObj.GetBool("$.boolValue"), false)
	assert.Equal(t, jsonObj.GetBool("$.objectValue.field1"), true)
	assert.Equal(t, jsonObj.GetInt("$.objectValue.field2Struct.f21"), 43)

	//testEntity := TestEntity{}
	//_ = jsonObject.GetObject("$.objectValue", &testEntity)
	//
	//fmt.Println(testEntity)
}

type TestEntity struct {
	Field1       bool
	Field2Struct TestEntity2
}

type TestEntity2 struct {
	F21     int
	K2array []int
}

// 普通数组
//
//	{
//	   "data": {
//	       "values": [
//	           {
//	               "name": "zhou",
//	               "age": 1
//	           },
//	           {
//	               "name": "song",
//	               "age": 2
//	           }
//	       ]
//	   }
//	}
func TestGet5(t *testing.T) {
	str := "{\"data\":{\"values\":[{\"name\":\"zhou\",\"age\":1},{\"name\":\"song\",\"age\":2}]}}"
	jsonObj, err := jsonobj.Load(str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	assert.Equal(t, jsonObj.GetString("$.data.values[0].name"), "zhou")
	assert.Equal(t, jsonObj.GetString("$.data.values[0].age"), "1")
	assert.Equal(t, jsonObj.GetString("$.data.values[1].name"), "song")
	assert.Equal(t, jsonObj.GetString("$.data.values[1].age"), "2")
}

// 二维数组
func TestGet5_1(t *testing.T) {
	str := "{\"data\":{\"values\":[[{\"name\":\"zhou\",\"age\":1},{\"name\":\"song\",\"age\":2}]]}}"
	jsonObject, err := jsonobj.Load(str)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	assert.Equal(t, jsonObject.GetString("$.data.values[0][0].name"), "zhou")
	assert.Equal(t, jsonObject.GetString("$.data.values[0][0].age"), "1")
	assert.Equal(t, jsonObject.GetString("$.data.values[0][1].name"), "song")
	assert.Equal(t, jsonObject.GetString("$.data.values[0][1].age"), "2")
}

//	{
//	   "intValue": 12,
//	   "intValue8": 12,
//	   "intValue16": 12,
//	   "intValue32": 12,
//	   "intValue64": 12,
//	   "stringValue": "haode",
//	   "boolValue": false,
//	   "objectValue": {
//	       "field1": true,
//	       "field2Struct": {
//	           "f21": 43,
//	           "k2array": [
//	               1,
//	               2,
//	               3,
//	               4
//	           ]
//	       }
//	   }
//	}
func TestRoot(t *testing.T) {
	str = "{\"intValue\":12,\"intValue8\":12,\"intValue16\":12,\"intValue32\":12,\"intValue64\":12,\"stringValue\":\"haode\",\"boolValue\":false,\"objectValue\":{\"field1\":true,\"field2Struct\":{\"f21\":43,\"k2array\":[1,2,3,4]}}}"
	obj, _ := jsonobj.Load(str)

	fmt.Println(util.ToJsonString(obj.Get("$")))
}
