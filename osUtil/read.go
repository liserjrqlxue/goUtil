package osUtil

import (
	"bufio"
	"bytes"
	"io"
	"io/fs"
	"regexp"

	"github.com/liserjrqlxue/goUtil/scannerUtil"
)

func FS2Array(file fs.File) []string {
	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scanner2Array(scanner)
}

// FS2MapArray 从 fs.File 读取到 []map[string]string
func FS2MapArray(file fs.File, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	var scanner = bufio.NewScanner(file)
	return scannerUtil.Scanner2MapArray(scanner, sep, skip)
}

// Bytes2MapArray 从字节切片读取到 []map[string]string
func Bytes2MapArray(data []byte, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)
	return scannerUtil.Scanner2MapArray(scanner, sep, skip)
}

// Reader2MapArray 从 io.Reader 读取到 []map[string]string
func Reader2MapArray(reader io.Reader, sep string, skip *regexp.Regexp) ([]map[string]string, []string) {
	scanner := bufio.NewScanner(reader)
	return scannerUtil.Scanner2MapArray(scanner, sep, skip)
}
