package textUtil

import (
	"bufio"
	"os"

	"github.com/liserjrqlxue/goUtil/scannerUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

// read file to []string, each line as item
func File2Array(fileName string) []string {
	var file, err = os.Open(fileName)
	simpleUtil.CheckErr(err)
	defer simpleUtil.DeferClose(file)

	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scanner2Array(scanner)
}

// read file to map[string]string, each line split by sep, first item as key and second item as value
func File2Map(fileName, sep string, override bool) (map[string]string, error) {
	var file, err = os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer simpleUtil.DeferClose(file)

	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scanner2Map(scanner, sep, override)
}
