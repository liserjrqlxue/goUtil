package xlsxUtil

import (
	"github.com/tealeg/xlsx/v3"

	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

type File struct {
	File *xlsx.File
}

func NewFile() *File {
	return &File{xlsx.NewFile()}
}
func (excel *File) Save(path string) {
	simpleUtil.CheckErr(excel.File.Save(path))
}

func OpenFile(fileName string) *File {
	var f, err = xlsx.OpenFile(fileName)
	simpleUtil.CheckErr(err)
	return &File{f}
}

func (excel *File) AddSheet(sheetName string) *xlsx.Sheet {
	var sheet, err = excel.File.AddSheet(sheetName)
	simpleUtil.CheckErr(err)
	return sheet
}

func (excel *File) AppendSheet(sheet xlsx.Sheet, sheetName string) *xlsx.Sheet {
	var newSheet, err = excel.File.AppendSheet(sheet, sheetName)
	simpleUtil.CheckErr(err)
	return newSheet
}

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

func GetRowArray(idx int, sheet *xlsx.Sheet) (rowArray []string) {
	var row, err = sheet.Row(idx)
	simpleUtil.CheckErr(err)
	for i := 0; i < sheet.MaxCol; i++ {
		var cell = row.GetCell(i)
		rowArray = append(rowArray, cell.Value)
	}
	return
}
