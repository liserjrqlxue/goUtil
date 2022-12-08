package stringsUtil

import (
	"strconv"
	"strings"

	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

// try to convert string to given width
func FormatWidth(width int, str string, fill byte) string {
	for len(str) < width {
		str = str + string(fill)
	}
	for len(str) > width {
		strs := strings.Split(str, "")
		if len(strs) > width {
			str = strings.Join(strs[0:width-1], "")
		} else {
			str = strings.Join(strs[0:len(strs)-1], "")
		}
	}
	return str
}

// plus for int in string format
func StringPlus(str string, num int) string {
	var strInt, err = strconv.Atoi(str)
	simpleUtil.CheckErr(err)
	strInt += num
	return strconv.Itoa(strInt)
}

// wrap of strconv.Atoi
func Atoi(str ...string) int {
	var v, e = strconv.Atoi(str[0])
	simpleUtil.CheckErr(e, str...)
	return v
}

// warp of strconv.ParseFloat
func Atof(str ...string) float64 {
	var v, e = strconv.ParseFloat(str[0], 32)
	simpleUtil.CheckErr(e, str...)
	return v
}

func Str2MapArray(str, sep1, sep2 string) (ma []map[string]string, title []string) {
	for i, s := range strings.Split(str, sep1) {
		if i == 0 {
			title = strings.Split(s, sep2)
		} else {
			var array = strings.Split(s, sep2)
			var item = make(map[string]string)
			for j, k := range array {
				item[title[j]] = k
			}
			ma = append(ma, item)
		}

	}
	return
}
