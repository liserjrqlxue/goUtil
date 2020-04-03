package scannerUtil

import (
	"bufio"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

func Scanner2Array(scanner *bufio.Scanner) []string {
	var array []string
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	simpleUtil.CheckErr(scanner.Err())
	return array
}
