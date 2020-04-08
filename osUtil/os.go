package osUtil

import (
	"os"

	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

func Create(fileName string) *os.File {
	var file, err = os.Create(fileName)
	simpleUtil.CheckErr(err)
	return file
}

func Write(file *os.File, b []byte) int {
	var n, err = file.Write(b)
	simpleUtil.CheckErr(err)
	return n
}

func WriteClose(file *os.File, b []byte) int {
	defer simpleUtil.DeferClose(file)
	var n, err = file.Write(b)
	simpleUtil.CheckErr(err)
	return n
}