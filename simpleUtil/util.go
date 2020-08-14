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
