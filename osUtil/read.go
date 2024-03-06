package osUtil

import (
	"bufio"
	"io/fs"
	"regexp"

	"github.com/liserjrqlxue/goUtil/scannerUtil"
)

func FS2Array(file fs.File) []string {
	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scanner2Array(scanner)
}

func FS2MapArray(file fs.File, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scanner2MapArray(scanner, sep, skip)
}
