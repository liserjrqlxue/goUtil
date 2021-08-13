package simpleUtil

import (
	"log"
	"regexp"
	"strings"
)

// CheckErr handle error
func CheckErr(err error, msg ...string) {
	if err != nil {
		//panic(err)
		log.Fatal(err, msg)
	}
}

type handle interface {
	Close() error
}

// DeferClose handle error while defer Close()
func DeferClose(h handle) {
	err := h.Close()
	CheckErr(err)
}

func HandleError(a interface{}, err error) interface{} {
	CheckErr(err)
	return a
}

func Slice2MapArray(slice [][]string) (db []map[string]string, title []string) {
	for i, array := range slice {
		if i == 0 {
			title = array
		} else {
			var item = make(map[string]string)
			for j := range array {
				item[title[j]] = array[j]
			}
			db = append(db, item)
		}
	}
	return
}

func JoinValue(item map[string]string, keys []string, sep string) string {
	var array []string
	for _, key := range keys {
		array = append(array, item[key])
	}
	return strings.Join(array, sep)
}

func Slice2MapMapArray(slice [][]string, keys ...string) (db map[string]map[string]string, title []string) {
	db = make(map[string]map[string]string)
	for i, array := range slice {
		if i == 0 {
			title = array
			for _, key := range keys {
				if !IsArrayContain(title, key) {
					panic("keys[" + key + "] not contain!")
				}
			}
		} else {
			var item = make(map[string]string)
			for j := range array {
				item[title[j]] = array[j]
			}
			var mainKey = JoinValue(item, keys, "\t")
			db[mainKey] = item
		}
	}
	return
}

func Slice2MapMapArrayMerge(slice [][]string, key, sep string) (db map[string]map[string]string, title []string) {
	var sepRegexp = regexp.MustCompile(sep)
	db = make(map[string]map[string]string)
	for i, array := range slice {
		if i == 0 {
			title = array
			if !IsArrayContain(title, key) {
				panic("key[" + key + "] not contain!")
			}
		} else {
			var item = make(map[string]string)
			for j, v := range array {
				if sepRegexp.MatchString(v) {
					log.Printf("WARN:\t[%d,%d]:[%s] contain sep[%s]\n", i, j, v, sep)
				}
				item[title[j]] = v
			}
			var mainKey = item[key]
			var mainItem, ok = db[mainKey]
			if ok {
				for k := range mainItem {
					mainItem[k] += sep + item[k]
				}
			} else {
				mainItem = item
			}
			db[mainKey] = mainItem
		}
	}
	return
}

// skip some WARN
func Slice2MapMapArrayMerge1(slice [][]string, key, sep string, skip map[int]bool) (db map[string]map[string]string, title []string) {
	var sepRegexp = regexp.MustCompile(sep)
	db = make(map[string]map[string]string)
	for i, array := range slice {
		if i == 0 {
			title = array
			if !IsArrayContain(title, key) {
				panic("key[" + key + "] not contain!")
			}
			var skips []string
			for index := range skip {
				skips = append(skips, title[index])
			}
			log.Printf("Skip merge warn of %+v", skips)
		} else {
			var item = make(map[string]string)
			for j, v := range array {
				if sepRegexp.MatchString(v) {
					if !skip[j] {
						var sepStr = sep
						if sep == "\n" {
							sepStr = `\n`
						}
						log.Printf("WARN:\t[%d,%d]:[%s] contain sep[%s]\n", i, j, v, sepStr)
					}
				}
				item[title[j]] = v
			}
			var mainKey = item[key]
			var mainItem, ok = db[mainKey]
			if ok {
				for k := range mainItem {
					mainItem[k] += sep + item[k]
				}
			} else {
				mainItem = item
			}
			db[mainKey] = mainItem
		}
	}
	return
}

func IsArrayContain(array []string, key string) bool {
	for _, item := range array {
		if item == key {
			return true
		}
	}
	return false
}
