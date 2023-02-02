package osUtil

import (
	"io"
	"log"
	"os"

	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

func Hostname() string {
	return simpleUtil.HandleError(os.Hostname()).(string)
}

func Create(fileName string) (file *os.File) {
	switch fileName {
	case "-":
		file = os.Stdout
	default:
		var err error
		file, err = os.Create(fileName)
		simpleUtil.CheckErr(err)
	}
	return
}

func Write(file *os.File, b []byte) int {
	var n, err = file.Write(b)
	simpleUtil.CheckErr(err)
	return n
}

func WriteClose(file *os.File, b []byte) int {
	defer simpleUtil.DeferClose(file)
	var n, err = file.Write(b)
	simpleUtil.CheckErr(err)
	return n
}

func Open(fileName string) *os.File {
	var file, err = os.Open(fileName)
	simpleUtil.CheckErr(err)
	return file
}

// FileExists check if a file exists and is not a directory
func FileExists(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// check if a file is exists and empty
func FileEmpty(fileName string) bool {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return false
	}
	if info.Size() == 0 {
		return true
	}
	return false
}

func CopyFile(dst, src string) (err error) {
	r, err := os.Open(src)
	if err != nil {
		return
	}
	defer simpleUtil.DeferClose(r)

	w, err := os.Create(dst)
	if err != nil {
		return
	}
	defer simpleUtil.DeferClose(w)

	n, err := io.Copy(w, r)
	if err != nil {
		return
	}
	log.Printf("CopyFile %d bytes[%s -> %s]", n, src, dst)
	return w.Sync()
}

func Symlink(source, dest string) error {
	_, err := os.Stat(dest)
	if err == nil {
		readLink, err := os.Readlink(dest)
		if err != nil {
			log.Printf("%v\n", err)
		}
		if readLink != source {
			log.Printf("dest is not symlink of source:[%s]->[%s]vs[%s]\n", dest, readLink, source)
			err = os.Symlink(source, dest)
			if err != nil {
				log.Printf("%v\n", err)
			}
		} else {
			log.Printf("dest is symlink of source:[%s]->[%s]", dest, readLink)
		}
	} else if os.IsNotExist(err) {
		err = os.Symlink(source, dest)
		if err != nil {
			log.Printf("Error: Symlink[%s->%s] err:%v", source, dest, err)
		}
	} else {
		log.Printf("Error: dest[%s] stat err:%v", dest, err)
	}
	return nil
}
