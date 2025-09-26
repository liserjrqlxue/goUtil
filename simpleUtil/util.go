package simpleUtil

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strings"
)

// CheckErr handle error
func CheckErr(err error, msg ...string) {
	if err != nil {
		// 构建完整的错误消息
		fullMsg := err.Error()
		if len(msg) > 0 {
			fullMsg = fmt.Sprintf("%s: %s", strings.Join(msg, " "), fullMsg)
		}

		// 使用 log.Printf 进行格式化输出，包含时间戳
		log.Printf("FATAL: %s", fullMsg)

		// 触发 panic
		panic(errors.New(fullMsg))
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

func HandleError[T any](a T, err error) T {
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
	var sepRegexp = regexp.MustCompile(fmt.Sprint(regexp.QuoteMeta(sep)))
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

// 通用浅拷贝函数
func ShallowCopy(src any) any {
	if src == nil {
		return nil
	}

	srcVal := reflect.ValueOf(src)

	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}

	if srcVal.Kind() != reflect.Struct {
		panic("ShallowCopy: src must be a struct or pointer to struct")
	}

	srcType := reflect.TypeOf(src)

	// 创建一个新的实例
	copyVal := reflect.New(srcType.Elem()).Elem()

	// 复制每个字段
	for i := 0; i < srcVal.Elem().NumField(); i++ {
		field := srcVal.Elem().Field(i)
		copyVal.Field(i).Set(field)
	}

	return copyVal.Addr().Interface()
}

func DeepCopyReflect(src any) any {
	srcVal := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	switch srcVal.Kind() {
	case reflect.Ptr:
		// 如果是指针，递归复制指针指向的值
		if srcVal.IsNil() {
			return nil
		}
		copyVal := reflect.New(srcType.Elem())
		copyVal.Elem().Set(reflect.ValueOf(DeepCopyReflect(srcVal.Elem().Interface())))
		return copyVal.Interface()
	case reflect.Struct:
		// 如果是结构体，逐字段复制
		copyVal := reflect.New(srcType).Elem()
		for i := 0; i < srcVal.NumField(); i++ {
			field := srcVal.Field(i)
			if field.CanSet() {
				copyVal.Field(i).Set(reflect.ValueOf(DeepCopyReflect(field.Interface())))
			}
		}
		return copyVal.Interface()
	case reflect.Slice:
		// 如果是切片，复制每个元素
		if srcVal.IsNil() {
			return nil
		}
		copyVal := reflect.MakeSlice(srcType, srcVal.Len(), srcVal.Cap())
		for i := 0; i < srcVal.Len(); i++ {
			copyVal.Index(i).Set(reflect.ValueOf(DeepCopyReflect(srcVal.Index(i).Interface())))
		}
		return copyVal.Interface()
	case reflect.Map:
		// 如果是映射，复制每个键值对
		if srcVal.IsNil() {
			return nil
		}
		copyVal := reflect.MakeMap(srcType)
		for _, key := range srcVal.MapKeys() {
			copyVal.SetMapIndex(key, reflect.ValueOf(DeepCopyReflect(srcVal.MapIndex(key).Interface())))
		}
		return copyVal.Interface()
	default:
		// 对于其他类型，直接返回值
		return src
	}
}
