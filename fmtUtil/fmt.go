package fmtUtil

import (
	"fmt"
	"io"
	"strings"

	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

func Base16toStr(base [16]byte) string {
	return fmt.Sprintf("%x", base)
}

// wrap of fmt.Fprint
func Fprint(w io.Writer, a ...interface{}) int {
	var n, err = fmt.Fprint(w, a...)
	simpleUtil.CheckErr(err)
	return n
}

// wrap of fmt.Fprintln
func Fprintln(w io.Writer, a ...interface{}) int {
	var n, err = fmt.Fprintln(w, a...)
	simpleUtil.CheckErr(err)
	return n
}

// wrap of fmt.Fprintf
func Fprintf(w io.Writer, format string, a ...interface{}) int {
	var n, err = fmt.Fprintf(w, format, a...)
	simpleUtil.CheckErr(err)
	return n
}

func FprintStringArray(w io.Writer, a []string, sep string) int {
	var n, err = fmt.Fprintln(w, strings.Join(a, sep))
	simpleUtil.CheckErr(err)
	return n
}
