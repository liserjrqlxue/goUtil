package cryptoUtil

import (
	"crypto/md5"

	"github.com/liserjrqlxue/goUtil/fmtUtil"
)

func Md5sum(str string) string {
	return fmtUtil.Base16toStr(md5.Sum([]byte(str)))
}
