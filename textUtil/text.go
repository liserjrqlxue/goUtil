package textUtil

import (
	"bufio"
	"log"
	"os"
	"regexp"

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
	return scannerUtil.Scan2Map(scanner, sep, override)
}

// read file to map[string]string and keys array, each line split by sep, first item as key and second item as value
func File2MapOrder(fileName, sep string, override bool) (map[string]string, []string, error) {
	var file, err = os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer simpleUtil.DeferClose(file)

	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scan2MapOrder(scanner, sep, override)
}

// read file to []map[string]string
func File2MapArray(fileName, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	file, err := os.Open(fileName)
	simpleUtil.CheckErr(err)
	defer simpleUtil.DeferClose(file)

	scanner := bufio.NewScanner(file)
	return scannerUtil.Scanner2MapArray(scanner, sep, skip)
}

// read files to []map[string]string
func Files2MapArray(fileNames []string, sep string, skip *regexp.Regexp) (Data []map[string]string, Title []string) {
	for _, fileName := range fileNames {
		data, title := File2MapArray(fileName, sep, skip)
		for _, item := range data {
			Data = append(Data, item)
		}
		if len(Title) == 0 {
			Title = title
		} else {
			if len(Title) != len(title) {
				log.Fatal("titles has different columns")
			} else {
				for i := 0; i < len(Title); i++ {
					if Title[i] != title[i] {
						log.Fatal("titles not equal")
					}
				}
			}
		}
	}
	return
}
