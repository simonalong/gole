package test

import (
	"github.com/simonalong/gole/maps"
	baseTime "github.com/simonalong/gole/time"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPutGet(t *testing.T) {
	dataMap := maps.New()
	dataMap.Put("a", 12)

	d, _ := dataMap.GetInt("a")
	assert.Equal(t, 12, d)
}

type PairTem struct {
	Left  interface{}
	Right interface{}
}

//type MultiPair struct {
//	Left   interface{}
//	Middle interface{}
//	Right  interface{}
//}

func TestSort1(t *testing.T) {
	dataMap := maps.New()
	dataMap.SetSort(true)
	dataMap.Put("a", 12)
	dataMap.Put("b", 13)
	dataMap.Put("c", 124)
	dataMap.Put("d", 54)
	dataMap.Put("e", 36)

	datas := []PairTem{
		{Left: "a", Right: 12},
		{Left: "b", Right: 13},
		{Left: "c", Right: 124},
		{Left: "d", Right: 54},
		{Left: "e", Right: 36},
	}

	// 循环
	for index, key := range dataMap.Keys() {
		val, _ := dataMap.Get(key)
		assert.Equal(t, datas[index].Left, key)
		assert.Equal(t, datas[index].Right, val)
	}
}

func TestSort2(t *testing.T) {
	dataMap := maps.New()
	dataMap.Put("a", 12)
	dataMap.Put("b", 13)
	dataMap.Put("c", 124)

	dataMap.SetSort(true)
	dataMap.Put("d", 54)
	dataMap.Put("e", 36)

	datas := []PairTem{
		{Left: "a", Right: 12},
		{Left: "b", Right: 13},
		{Left: "c", Right: 124},
		{Left: "d", Right: 54},
		{Left: "e", Right: 36},
	}

	// 循环
	for index, key := range dataMap.Keys() {
		val, _ := dataMap.Get(key)
		if index < 3 {
			continue
		}
		assert.Equal(t, datas[index].Left, key)
		assert.Equal(t, datas[index].Right, val)
	}
}

func TestSort3(t *testing.T) {
	dataMap := maps.New()
	dataMap.SetSort(true)
	dataMap.Put("a", 12)
	dataMap.Put("b", 13)
	dataMap.Put("c", 124)

	dataMap.SetSort(false)
	dataMap.Put("d", 54)
	dataMap.Put("e", 36)

	// 循环：请使用keys循环
	//for _, key := range dataMap.Keys() {
	//	val, _ := dataMap.Get(key)
	//	fmt.Println(key, val)
	//}
}

func TestFromEntity1(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time `column:"ts"`
		Name    string    `column:"name"`
		Age     int       `column:"age"`
		Address string    `column:"address"`
	}

	entity1 := DemoEntity{
		Ts:      time.Now(),
		Name:    "test",
		Age:     22,
		Address: "浙江",
	}

	dataMap := maps.FromEntity(entity1)
	//fmt.Println(dataMap.ToString())

	for index, key := range dataMap.Keys() {
		switch index {
		case 0:
			d, _ := dataMap.GetTime(key)
			assert.Equal(t, baseTime.TimeToStringYmdHmsS(entity1.Ts), baseTime.TimeToStringYmdHmsS(d))
		case 1:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, entity1.Name, d)
		case 2:
			d, _ := dataMap.GetInt(key)
			assert.Equal(t, entity1.Age, d)
		case 3:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, entity1.Address, d)
		}
	}
}

func TestFromEntity2(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time `json:"ts"`
		Name    string    `json:"name"`
		Age     int       `json:"age"`
		Address string    `json:"address"`
	}

	entity2 := DemoEntity{
		Ts:      time.Now(),
		Name:    "test",
		Age:     22,
		Address: "浙江",
	}

	dataMap := maps.FromEntity(entity2)

	for index, key := range dataMap.Keys() {
		switch index {
		case 0:
			d, _ := dataMap.GetTime(key)
			assert.Equal(t, baseTime.TimeToStringYmdHmsS(entity2.Ts), baseTime.TimeToStringYmdHmsS(d))
		case 1:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, entity2.Name, d)
		case 2:
			d, _ := dataMap.GetInt(key)
			assert.Equal(t, entity2.Age, d)
		case 3:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, entity2.Address, d)
		}
	}
}

func TestFromEntity3(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time
		Name    string
		Age     int
		Address string
	}

	entity3 := DemoEntity{
		Ts:      time.Now(),
		Name:    "test",
		Age:     22,
		Address: "浙江",
	}
	dataMap := maps.FromEntity(entity3)

	for index, key := range dataMap.Keys() {
		switch index {
		case 0:
			d, _ := dataMap.GetTime(key)
			assert.Equal(t, baseTime.TimeToStringYmdHmsS(entity3.Ts), baseTime.TimeToStringYmdHmsS(d))
		case 1:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, entity3.Name, d)
		case 2:
			d, _ := dataMap.GetInt(key)
			assert.Equal(t, entity3.Age, d)
		case 3:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, entity3.Address, d)
		}
	}
}

func TestFromEntity4(t *testing.T) {
	current := time.Now()
	baseMap := maps.OfSort(
		"Ts", current,
		"Name", "test",
		"Age", 22,
		"Address", "浙江",
	)

	dataMap := maps.FromEntity(baseMap)

	for index, key := range dataMap.Keys() {
		switch index {
		case 0:
			d, _ := dataMap.GetTime(key)
			assert.Equal(t, baseTime.TimeToStringYmdHmsS(current), baseTime.TimeToStringYmdHmsS(d))
		case 1:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, "test", d)
		case 2:
			d, _ := dataMap.GetInt(key)
			assert.Equal(t, 22, d)
		case 3:
			d, _ := dataMap.GetString(key)
			assert.Equal(t, "浙江", d)
		}
	}
}

func TestFromEntity5(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time
		Name    string
		Age     int
		Address string
	}

	type DemoEntity2 struct {
		Name string
		Data *DemoEntity
	}
	entity3 := DemoEntity2{
		Name: "test",
		Data: &DemoEntity{
			Ts:      time.Now(),
			Name:    "test",
			Age:     22,
			Address: "浙江",
		},
	}
	_ = maps.FromEntity(entity3)
	//fmt.Println(dataMap.ToJsonOfSort())
}

func TestFromEntity6(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time
		Name    string
		Age     int
		Address string
	}

	type DemoEntity2 struct {
		Name  string
		Data1 []string
		Data  []*DemoEntity
	}
	entity3 := DemoEntity2{
		Name:  "test",
		Data1: []string{"a", "b", "c"},
		Data: []*DemoEntity{{
			Ts:      time.Now(),
			Name:    "test",
			Age:     21,
			Address: "浙江",
		}, {
			Ts:      time.Now(),
			Name:    "test",
			Age:     22,
			Address: "浙江",
		}},
	}
	_ = maps.FromEntity(entity3)
	//fmt.Println(dataMap.ToJsonOfSort())
}

func TestFromMap(t *testing.T) {
	dataMap := map[string]interface{}{
		"a": 12,
		"b": 122,
		"c": 42,
		"d": 57,
	}
	OrmMap := maps.FromMap(dataMap)
	if OrmMap == nil {
		return
	}

	d, _ := OrmMap.GetInt("a")
	assert.Equal(t, dataMap["a"], d)

	d, _ = OrmMap.GetInt("b")
	assert.Equal(t, dataMap["b"], d)
	d, _ = OrmMap.GetInt("c")
	assert.Equal(t, dataMap["c"], d)
	d, _ = OrmMap.GetInt("d")
	assert.Equal(t, dataMap["d"], d)
}

func TestFromJson1(t *testing.T) {
	jsonData := "{\"a\":12,\"b\":32,\"d\":{\"e\":12}}"
	OrmMap, _ := maps.FromJson(jsonData)
	d, _ := OrmMap.GetInt("a")
	assert.Equal(t, 12, d)
	d, _ = OrmMap.GetInt("b")
	assert.Equal(t, 32, d)
}

func TestFrom1(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time
		Name    string
		Age     int
		Address string
	}

	entity3 := DemoEntity{
		Ts:      time.Now(),
		Name:    "test",
		Age:     22,
		Address: "浙江",
	}
	dataMap, _ := maps.From(entity3)
	val, _ := dataMap.GetString("name")
	assert.Equal(t, val, "test")

	map1 := map[string]interface{}{
		"a":    12,
		"name": "zhou",
		"c":    42,
		"d":    57,
	}
	dataMap2, _ := maps.From(map1)
	val1, _ := dataMap2.GetString("name")
	assert.Equal(t, "zhou", val1)

	jsonData := "{\"a\":12,\"b\":32,\"d\":{\"e\":12}}"
	dataMap3, _ := maps.From(jsonData)
	val2, _ := dataMap3.GetInt("a")
	assert.Equal(t, 12, val2)
}

func TestToJson(t *testing.T) {
	jsonData := "{\"a\":12,\"b\":32,\"d\":{\"e\":12}}"
	OrmMap, _ := maps.FromJson(jsonData)
	jsonContent := OrmMap.ToJson()
	assert.Equal(t, jsonData, jsonContent)
}

func TestClone1(t *testing.T) {
	dataMap := maps.NewSort()
	dataMap.Put("a", 12)
	dataMap.Put("b", 13)
	dataMap.Put("c", 124)

	var expect []PairTem
	var clone []PairTem
	cloneMap := dataMap.Clone()
	for _, key := range dataMap.Keys() {
		val, _ := dataMap.Get(key)
		expect = append(expect, PairTem{key, val})
	}

	for _, key := range cloneMap.Keys() {
		val, _ := cloneMap.Get(key)
		clone = append(clone, PairTem{key, val})
	}

	for index := range expect {
		assert.Equal(t, expect[index].Left, clone[index].Left)
		assert.Equal(t, expect[index].Right, clone[index].Right)
	}
}

func TestToEntity1(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time `column:"ts"`
		Name    string    `column:"name"`
		Age     int       `column:"age"`
		Address string    `column:"address"`
	}

	entity1 := DemoEntity{
		Ts:      time.Now(),
		Name:    "test",
		Age:     22,
		Address: "浙江",
	}
	entityMap := maps.FromEntity(entity1)

	entity1Expect := DemoEntity{}
	err := entityMap.ToEntity(&entity1Expect)

	assert.Equal(t, nil, err)
	assert.Equal(t, entity1Expect.Name, entity1.Name)
	assert.Equal(t, entity1Expect.Age, entity1.Age)
	assert.Equal(t, entity1Expect.Address, entity1.Address)
}

func TestToEntity2(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time `json:"ts"`
		Name    string    `json:"name"`
		Age     int       `json:"age"`
		Address string    `json:"address"`
	}

	entity1 := DemoEntity{
		Ts:      time.Now(),
		Name:    "test",
		Age:     22,
		Address: "浙江",
	}
	entityMap := maps.FromEntity(entity1)

	entity1Expect := DemoEntity{}
	err := entityMap.ToEntity(&entity1Expect)

	assert.Equal(t, nil, err)
	assert.Equal(t, entity1.Ts, entity1Expect.Ts)
	assert.Equal(t, entity1Expect.Name, entity1.Name)
	assert.Equal(t, entity1Expect.Age, entity1.Age)
	assert.Equal(t, entity1Expect.Address, entity1.Address)
}

func TestToEntity3(t *testing.T) {
	type DemoEntity struct {
		Ts      time.Time
		Name    string
		Age     int
		Address string
	}

	entity1 := DemoEntity{
		Ts:      time.Now(),
		Name:    "test",
		Age:     22,
		Address: "浙江",
	}
	entityMap := maps.FromEntity(entity1)

	entity1Expect := DemoEntity{}
	err := entityMap.ToEntity(&entity1Expect)

	assert.Equal(t, nil, err)
	assert.Equal(t, entity1.Ts, entity1Expect.Ts)
	assert.Equal(t, entity1Expect.Name, entity1.Name)
	assert.Equal(t, entity1Expect.Age, entity1.Age)
	assert.Equal(t, entity1Expect.Address, entity1.Address)
}

func TestToString1(t *testing.T) {
	dataMap := maps.New()
	dataMap.Put("a", 12)
	dataMap.Put("b", 13)
	dataMap.Put("c", 124)

	//fmt.Print(dataMap.ToString())
}

func TestToString2(t *testing.T) {
	dataMap := maps.NewSort()
	dataMap.Put("a", 12)
	dataMap.Put("b", 13)
	dataMap.Put("c", 124)

	assert.Equal(t, "[\"a\":12,\"b\":13,\"c\":124]", dataMap.ToString())
}

func TestOf1(t *testing.T) {
	pMap := maps.Of("a", 1, "b", 12)
	v, _ := pMap.GetInt("a")
	assert.Equal(t, 1, v)
}

func TestOf2(t *testing.T) {
	pMap := maps.Of("a", 1, "b", 12, "k3", 13, "k4", 14)
	v, _ := pMap.GetInt("a")
	assert.Equal(t, 1, v)
}

func TestGetBaseMap(t *testing.T) {
	pMap := maps.Of("a", maps.Of("aa", 12))
	v, _ := pMap.GetMaps("a")

	d, _ := v.GetString("aa")
	assert.Equal(t, d, "12")
}

//func TestToJsonOfSort1(t *testing.T) {
//	baseMap := util.OfSort("a", 12, "b", 2, "c", 3)
//	assert.Equal(t, "{\"a\":12,\"b\":2,\"c\":3}", baseMap.ToJsonOfSort())
//
//	baseMap = util.OfSort("c", 12, "b", 2, "e", 3)
//	assert.Equal(t, "{\"c\":12,\"b\":2,\"e\":3}", baseMap.ToJsonOfSort())
//
//	timeData, _ := baseTime.ParseTime("2024-08-28")
//
//	baseMap = util.OfSort("c", true, "b", false, "e", "test", "f", timeData, "g")
//	assert.Equal(t, "{\"c\":true,\"b\":false,\"e\":\"test\",\"f\":\"2024-08-28 00:00:00 +0000 UTC\"}", baseMap.ToJsonOfSort())
//}

func TestToJsonOfSort2(t *testing.T) {
	type Entity struct {
		Name string
		Age  int
		Man  bool
		Addr string
	}

	entity := Entity{
		Name: "test",
		Age:  22,
		Man:  true,
		Addr: "hang",
	}
	baseMap := maps.OfSort("c", 12, "b", entity, "e", 3)
	assert.Equal(t, "{\"c\":12,\"b\":{\"name\":\"test\",\"age\":22,\"man\":true,\"addr\":\"hang\"},\"e\":3}", baseMap.ToJsonOfSort())

	baseMap = maps.OfSort("c", 12, "b", &entity, "e", 3)
	assert.Equal(t, "{\"c\":12,\"b\":{\"name\":\"test\",\"age\":22,\"man\":true,\"addr\":\"hang\"},\"e\":3}", baseMap.ToJsonOfSort())
}

func TestDeepMap(t *testing.T) {
	dataMap := maps.Of("a", maps.Of("b", maps.Of("c", 12)))
	data, have := dataMap.AsDeepMap().GetInt("a.b.c")
	assert.Equal(t, true, have)
	assert.Equal(t, 12, data)
}
