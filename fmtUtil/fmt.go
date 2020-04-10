package fmtUtil

import "fmt"

func Base16toStr(base [16]byte) string {
	return fmt.Sprintf("%x", base)
}
