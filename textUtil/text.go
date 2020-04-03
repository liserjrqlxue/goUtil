package textUtil

import (
	"bufio"
	"github.com/liserjrqlxue/goUtil/scannerUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
	"os"
)

func File2Array(fileName string) []string {
	var file, err = os.Open(fileName)
	simpleUtil.CheckErr(err)
	defer simpleUtil.DeferClose(file)

	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scanner2Array(scanner)
}
