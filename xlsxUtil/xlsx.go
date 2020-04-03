package xlsxUtil

import "github.com/liserjrqlxue/goUtil/simpleUtil"
import "github.com/tealeg/xlsx/v2"

func AddSheet(excel *xlsx.File, sheetName string) *xlsx.Sheet {
	var sheet, err = excel.AddSheet(sheetName)
	simpleUtil.CheckErr(err)
	return sheet
}

func AddSheets(excel *xlsx.File, sheetNames []string) {
	for _, sheetName := range sheetNames {
		AddSheet(excel, sheetName)
	}
}

func AddArray2Row(rows []string, row *xlsx.Row) {
	for _, str := range rows {
		row.AddCell().SetString(str)
	}
}

func AddSlice2Sheet(slice [][]string, sheet *xlsx.Sheet) {
	for _, array := range slice {
		AddArray2Row(array, sheet.AddRow())
	}
}

func AddMap2Row(item map[string]string, keys []string, row *xlsx.Row) {
	for _, key := range keys {
		row.AddCell().SetString(item[key])
	}
}

func AddMapArray2Sheet(array []map[string]string, keys []string, sheet *xlsx.Sheet) {
	for _, item := range array {
		AddMap2Row(item, keys, sheet.AddRow())
	}
}
