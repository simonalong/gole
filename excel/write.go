package excel

import (
	"errors"
	"fmt"
	"github.com/simonalong/gole/logger"
	"github.com/simonalong/gole/maps"
	"github.com/simonalong/gole/util"
	"github.com/tealeg/xlsx"
)

func AddMaps(targetFile string, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	defaultSheet := excelFile.Sheets[0]
	var keys []string
	for k, _ := range dataMaps[0] {
		keys = append(keys, k)
	}
	for _, dataMap := range dataMaps {
		row := defaultSheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AddEntities[T any](targetFile string, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AddMaps(targetFile, dataMaps)
}

func AddDatas(targetFile string, datasList [][]string) error {
	if len(datasList) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	defaultSheet := excelFile.Sheets[0]
	for _, datas := range datasList {
		row := defaultSheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AddSheetEntities[T any](targetFile, sheetName string, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AddSheetMaps(targetFile, sheetName, dataMaps)
}

func AddSheetMaps(targetFile, sheetName string, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet, err := excelFile.AddSheet(sheetName)
	if err != nil {
		logger.Errorf("创建sheet【%v】异常：%v", sheetName, err.Error())
		return err
	}
	var keys []string
	for k, _ := range dataMaps[0] {
		keys = append(keys, k)
	}
	for _, dataMap := range dataMaps {
		row := sheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AddSheetDatas(targetFile, sheetName string, datasList [][]string) error {
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet, err := excelFile.AddSheet(sheetName)
	if err != nil {
		logger.Errorf("创建sheet【%v】异常：%v", sheetName, err.Error())
		return err
	}

	for _, datas := range datasList {
		row := sheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AppendDatas(targetFile string, datasList [][]string) error {
	if len(datasList) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheets[0]
	for _, datas := range datasList {
		row := sheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AppendMaps(targetFile string, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheets[0]
	var keys []string
	for k, _ := range dataMaps[0] {
		keys = append(keys, k)
	}
	for _, dataMap := range dataMaps {
		row := sheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AppendEntities[T any](targetFile string, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AppendMaps(targetFile, dataMaps)
}

func AppendSheetDatas(targetFile, sheetName string, datasList [][]string) error {
	if len(datasList) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheet[sheetName]
	if sheet == nil {
		logger.Errorf("sheet【%v】不存在", sheetName)
		return errors.New(fmt.Sprintf("sheet【%v】不存在", sheetName))
	}
	for _, datas := range datasList {
		row := sheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}
func AppendSheetMaps(targetFile, sheetName string, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheet[sheetName]
	if sheet == nil {
		logger.Errorf("sheet【%v】不存在", sheetName)
		return errors.New(fmt.Sprintf("sheet【%v】不存在", sheetName))
	}
	var keys []string
	for k, _ := range dataMaps[0] {
		keys = append(keys, k)
	}
	for _, dataMap := range dataMaps {
		row := sheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}
func AppendSheetEntities[T any](targetFile, sheetName string, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AppendSheetMaps(targetFile, sheetName, dataMaps)
}

func AddMapsWithBanner(targetFile string, banner *maps.GoleMap, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	defaultSheet := excelFile.Sheets[0]
	keys := banner.Keys()
	bannerRow := defaultSheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}

	for _, dataMap := range dataMaps {
		row := defaultSheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AddEntitiesWithBanner[T any](targetFile string, banner *maps.GoleMap, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AddMapsWithBanner(targetFile, banner, dataMaps)
}

func AddDatasWithBanner(targetFile string, banner *maps.GoleMap, datasList [][]string) error {
	if len(datasList) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	defaultSheet := excelFile.Sheets[0]
	bannerRow := defaultSheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}

	for _, datas := range datasList {
		row := defaultSheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AddSheetEntitiesWithBanner[T any](targetFile, sheetName string, banner *maps.GoleMap, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AddSheetMapsWithBanner(targetFile, sheetName, banner, dataMaps)
}

func AddSheetMapsWithBanner(targetFile, sheetName string, banner *maps.GoleMap, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet, err := excelFile.AddSheet(sheetName)
	if err != nil {
		logger.Errorf("创建sheet【%v】异常：%v", sheetName, err.Error())
		return err
	}
	keys := banner.Keys()
	bannerRow := sheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}
	for _, dataMap := range dataMaps {
		row := sheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AddSheetDatasWithBanner(targetFile, sheetName string, banner *maps.GoleMap, datasList [][]string) error {
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet, err := excelFile.AddSheet(sheetName)
	if err != nil {
		logger.Errorf("创建sheet【%v】异常：%v", sheetName, err.Error())
		return err
	}

	bannerRow := sheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}
	for _, datas := range datasList {
		row := sheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AppendDatasWithBanner(targetFile string, banner *maps.GoleMap, datasList [][]string) error {
	if len(datasList) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheets[0]
	bannerRow := sheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}
	for _, datas := range datasList {
		row := sheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AppendMapsWithBanner(targetFile string, banner *maps.GoleMap, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheets[0]
	keys := banner.Keys()
	bannerRow := sheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}
	for _, dataMap := range dataMaps {
		row := sheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}

func AppendEntitiesWithBanner[T any](targetFile string, banner *maps.GoleMap, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AppendMapsWithBanner(targetFile, banner, dataMaps)
}

func AppendSheetDatasWithBanner(targetFile, sheetName string, banner *maps.GoleMap, datasList [][]string) error {
	if len(datasList) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheet[sheetName]
	if sheet == nil {
		logger.Errorf("sheet【%v】不存在", sheetName)
		return errors.New(fmt.Sprintf("sheet【%v】不存在", sheetName))
	}
	bannerRow := sheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}
	for _, datas := range datasList {
		row := sheet.AddRow()
		for _, data := range datas {
			cell := row.AddCell()
			cell.Value = data
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}
func AppendSheetMapsWithBanner(targetFile, sheetName string, banner *maps.GoleMap, dataMaps []map[string]any) error {
	if len(dataMaps) == 0 {
		return nil
	}
	excelFile, err := xlsx.OpenFile(targetFile)
	if err != nil {
		logger.Errorf("读取文件失败，路径：%v，错误：%v", targetFile, err.Error())
		return err
	}
	sheet := excelFile.Sheet[sheetName]
	if sheet == nil {
		logger.Errorf("sheet【%v】不存在", sheetName)
		return errors.New(fmt.Sprintf("sheet【%v】不存在", sheetName))
	}
	keys := banner.Keys()
	bannerRow := sheet.AddRow()
	for _, key := range banner.Keys() {
		cell := bannerRow.AddCell()
		cell.Value, _ = banner.GetString(key)
	}
	for _, dataMap := range dataMaps {
		row := sheet.AddRow()
		for _, key := range keys {
			if val, ok := dataMap[key]; ok {
				cell := row.AddCell()
				cell.Value = util.ToString(val)
			}
		}
	}

	err = excelFile.Save(targetFile)
	if err != nil {
		logger.Errorf("保存excel文件 %v 失败：%v", targetFile, err)
		return err
	}
	return nil
}
func AppendSheetEntitiesWithBanner[T any](targetFile, sheetName string, banner *maps.GoleMap, entities []T) error {
	if len(entities) == 0 {
		return nil
	}
	var dataMaps []map[string]any
	for _, entity := range entities {
		dataMap, _ := util.JsonToMap(util.ToJsonString(entity))
		dataMaps = append(dataMaps, dataMap)
	}
	return AppendSheetMapsWithBanner(targetFile, sheetName, banner, dataMaps)
}
