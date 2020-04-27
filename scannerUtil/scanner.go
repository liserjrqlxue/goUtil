package scannerUtil

import (
	"errors"
	"regexp"
	"strings"

	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

type Scanner interface {
	Scan() bool
	Text() string
	Err() error
}

func Scanner2Array(scanner Scanner) []string {
	var array []string
	for scanner.Scan() {
		array = append(array, scanner.Text())
	}
	simpleUtil.CheckErr(scanner.Err())
	return array
}

func Scan2Map(scanner Scanner, sep string, override bool) (db map[string]string, err error) {
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
		err = sErr
	}
	return
}

func Scan2MapOrder(scanner Scanner, sep string, override bool) (db map[string]string, keys []string, err error) {
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
		keys = append(keys, array[0])
	}
	var sErr = scanner.Err()
	if sErr != nil {
		err = sErr
	}
	return
}

func ScanTitle(scanner Scanner, sep string, skip *regexp.Regexp) (title []string) {
	for scanner.Scan() {
		var line = scanner.Text()
		if skip != nil && skip.MatchString(line) {
			continue
		}
		title = strings.Split(line, sep)
		break
	}
	simpleUtil.CheckErr(scanner.Err())
	return
}

func Scanner2MapArray(scanner Scanner, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	var mapArray []map[string]string
	var title = ScanTitle(scanner, sep, skip)
	for scanner.Scan() {
		line := scanner.Text()
		if skip != nil && skip.MatchString(line) {
			continue
		}
		array := strings.Split(line, sep)
		var dataHash = make(map[string]string)
		for j, k := range array {
			dataHash[title[j]] = k
		}
		mapArray = append(mapArray, dataHash)
	}
	simpleUtil.CheckErr(scanner.Err())
	return mapArray, title
}
