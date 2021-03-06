package jsonUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	. "github.com/liserjrqlxue/goUtil/simpleUtil"
)

// warpper of json.MarshalIndent
func JsonIndent(v interface{}, prefix, indent string) (b []byte, err error) {
	b, err = json.MarshalIndent(v, prefix, indent)
	b = bytes.Replace(b, []byte("\\u003c"), []byte("<"), -1)
	b = bytes.Replace(b, []byte("\\u003e"), []byte(">"), -1)
	b = bytes.Replace(b, []byte("\\u0026"), []byte("&"), -1)
	return
}

func MarshalString(v interface{}) string {
	var b, err = json.Marshal(v)
	CheckErr(err)
	return string(b)
}

func MarshalIndentString(v interface{}, prefix, indent string) string {
	var b, err = json.MarshalIndent(v, prefix, indent)
	CheckErr(err)
	return string(b)
}

func Json2file(json []byte, filenName string) error {
	file, err := os.Create(filenName)
	if err != nil {
		return err
	}
	defer DeferClose(file)

	c, err := file.Write(json)
	if err != nil {
		return err
	}
	fmt.Printf("write %d byte to %s\n", c, filenName)

	return nil
}

func Json2File(fileName string, a interface{}) error {
	b, err := json.MarshalIndent(a, "", "\t")
	if err != nil {
		return err
	}
	return Json2file(b, fileName)
}

func Json2rawFile(fileName string, a interface{}) error {
	b, err := JsonIndent(a, "", "\t")
	if err != nil {
		return err
	}
	return Json2file(b, fileName)
}

func Json2MapMap(jsonBlob []byte) map[string]map[string]string {
	var data = make(map[string]map[string]string)
	err := json.Unmarshal(jsonBlob, &data)
	CheckErr(err)
	return data
}

func Json2Map(jsonBlob []byte) map[string]string {
	var data = make(map[string]string)
	err := json.Unmarshal(jsonBlob, &data)
	CheckErr(err)
	return data
}

func Json2MapInt(jsonBlob []byte) map[string]int {
	var data = make(map[string]int)
	err := json.Unmarshal(jsonBlob, &data)
	CheckErr(err)
	return data
}

func Json2MapBool(jsonBlob []byte) map[string]bool {
	var data = make(map[string]bool)
	err := json.Unmarshal(jsonBlob, &data)
	CheckErr(err)
	return data
}

func JsonFile2MapMap(fileName string) map[string]map[string]string {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	logLoadJson(len(b), fileName)
	return Json2MapMap(b)
}

func JsonFile2Map(fileName string) map[string]string {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	logLoadJson(len(b), fileName)
	return Json2Map(b)
}

func JsonFile2MapInt(fileName string) map[string]int {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	logLoadJson(len(b), fileName)
	return Json2MapInt(b)
}

func JsonFile2MapBool(fileName string) map[string]bool {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	logLoadJson(len(b), fileName)
	return Json2MapBool(b)
}

func JsonFile2Interface(fileName string) interface{} {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	logLoadJson(len(b), fileName)
	var data interface{}
	err = json.Unmarshal(b, &data)
	CheckErr(err)
	return data
}

func JsonFile2Data(fileName string, v interface{}) {
	b, err := ioutil.ReadFile(fileName)
	CheckErr(err)
	logLoadJson(len(b), fileName)
	CheckErr(json.Unmarshal(b, v))
}

func logLoadJson(size int, fileName string) {
	log.Printf("load %10d byte from %s\n", size, fileName)
}
