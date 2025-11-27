package excel

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func ReadAllDatas(filePath string) ([][][]string, error) {
	return xlsx.FileToSlice(filePath)
}

// ReadDatas 默认读取第一页
func ReadDatas(filePath string) ([][]string, error) {
	allData, err := xlsx.FileToSlice(filePath)
	if err != nil {
		return nil, err
	}
	return allData[0], err
}

func ReadSheetDatas(filePath, sheetName string) ([][]string, error) {
	allData, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	val, ok := allData.Sheet[sheetName]
	if !ok {
		return nil, fmt.Errorf("sheet %s not found", sheetName)
	}
	var datasList [][]string
	for _, row := range val.Rows {
		var datas []string
		for _, cell := range row.Cells {
			datas = append(datas, cell.Value)
		}
		datasList = append(datasList, datas)
	}
	return datasList, nil
}
