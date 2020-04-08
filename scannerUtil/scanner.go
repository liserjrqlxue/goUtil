package scannerUtil

import (
	"bufio"
	"errors"
	"strings"

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

func Scanner2Map(scanner *bufio.Scanner, sep string, override bool) (db map[string]string, err error) {
	db = make(map[string]string)
	for scanner.Scan() {
		var line = scanner.Text()
		array := strings.Split(line, sep)
		array = append(array, "NA", "NA")
		var v, ok = db[array[0]]
		if ok && v != array[1] && !override {
			err = errors.New("dup key[" + array[0] + "],different value:[" + v + "]vs[" + array[1] + "]")
		}
		db[array[0]] = array[1]
	}
	var sErr = scanner.Err()
	if sErr != nil {
		return db, sErr
	} else {
		return
	}
}
