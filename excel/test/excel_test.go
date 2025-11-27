package test

import (
	"encoding/json"
	"fmt"
	"github.com/simonalong/gole/excel"
	"github.com/simonalong/gole/file"
	"github.com/simonalong/gole/maps"
	"testing"
)

func TestAddData(t *testing.T) {
	filePath := "./resources/test.xlsx"
	var datasList [][]string
	datasList = append(datasList, []string{"zhou1", "ok1"})
	datasList = append(datasList, []string{"zhou2", "ok2"})
	err := excel.AddDatas(filePath, datasList)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestAddSheetData(t *testing.T) {
	filePath := "./resources/test.xlsx"
	var datasList [][]string
	datasList = append(datasList, []string{"zhou1", "ok1"})
	datasList = append(datasList, []string{"zhou2", "ok2"})
	err := excel.AddSheetDatas(filePath, "demo2", datasList)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestAppendSheetData(t *testing.T) {
	//filePath := "./resources/test.xlsx"
	var datasList [][]string
	datasList = append(datasList, []string{"zhou1", "ok1"})
	datasList = append(datasList, []string{"zhou2", "ok2"})
	//err := AppendSheetData(filePath, "demo1", datasList)
	//if err != nil {
	//	t.Errorf("%v", err)
	//}
}

func TestAddMaps(t *testing.T) {
	var dataMaps = []map[string]string{}
	dataMap := make(map[string]string)
	dataMap["a1"] = "a1"
	dataMap["a2"] = "a2"
	dataMap["a3"] = "a3"
	dataMap["a4"] = "a4"
	dataMap["a5"] = "a5"
	dataMaps = append(dataMaps, dataMap)

	dataMap1 := make(map[string]string)
	dataMap1["a1"] = "b1"
	dataMap1["a2"] = "b2"
	dataMap1["a3"] = "b3"
	dataMap1["a4"] = "b4"
	dataMap1["a5"] = "b5"
	dataMaps = append(dataMaps, dataMap)

	dataMap2 := make(map[string]string)
	dataMap2["a1"] = "c1"
	dataMap2["a2"] = "c2"
	dataMap2["a3"] = "c3"
	dataMap2["a4"] = "c4"
	dataMap2["a5"] = "c5"
	dataMaps = append(dataMaps, dataMap)
	for _, dataMap := range dataMaps {
		for k, v := range dataMap {
			fmt.Println(k, "=", v)
		}
	}
}

func TestBanner2(t *testing.T) {
	filePath := "./resources/test.xlsx"
	bannerMap := maps.OfSort("deviceId", "设备id", "at", "时间", "value", "数据")
	dataStrs := file.ReadFile("./resources/data.json")

	dataMaps := map[string][]*OneData{}
	//util.StrToObject(dataStrs, &dataMaps)
	json.Unmarshal([]byte(dataStrs), &dataMaps)
	var dataList []*OneData
	for _, datas := range dataMaps {
		dataList = append(dataList, datas...)
	}
	excel.AddEntitiesWithBanner(filePath, bannerMap, dataList)
}

type OneData struct {
	DeviceId string `json:"deviceId"`
	At       string `json:"at"`
	Value    string `json:"value"`
}
