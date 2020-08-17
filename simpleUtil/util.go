package simpleUtil

import (
	"log"
)

// handle error
func CheckErr(err error, msg ...string) {
	if err != nil {
		//panic(err)
		log.Fatal(err, msg)
	}
}

type handle interface {
	Close() error
}

// handle error while defer Close()
func DeferClose(h handle) {
	err := h.Close()
	CheckErr(err)
}

func HandleError(a interface{}, err error) interface{} {
	CheckErr(err)
	return a
}

func Slice2MapArray(slice [][]string) (db []map[string]string) {
	var title []string
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

func Slice2MapMapArray(slice [][]string, key string) map[string]map[string]string {
	var db = make(map[string]map[string]string)
	var title []string
	for i, array := range slice {
		if i == 0 {
			title = array
			if !IsArrayContain(title, key) {
				panic("key[" + key + "] not contain!")
			}
		} else {
			var item = make(map[string]string)
			for j := range array {
				item[title[j]] = array[j]
			}
			db[item[key]] = item
		}
	}
	return db
}

func IsArrayContain(array []string, key string) bool {
	for _, item := range array {
		if item == key {
			return true
		}
	}
	return false
}
