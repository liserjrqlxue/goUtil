package cryptoUtil

import (
	"crypto/md5"
)

func Md5sum(str string) string {
	var t = md5.Sum([]byte(str))
	return string(t[:])

}
