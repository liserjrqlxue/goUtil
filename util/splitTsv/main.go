package main

import (
	"flag"
	"fmt"
	"github.com/liserjrqlxue/goUtil/fmtUtil"
	"github.com/liserjrqlxue/goUtil/osUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
	"github.com/liserjrqlxue/goUtil/textUtil"
	"log"
	"os"
)

var (
	in = flag.String(
		"i",
		"",
		"input tsv with header",
	)
	prefix = flag.String(
		"p",
		"",
		"output prefix, split input to prefix.[key].tsv, default is same to input",
	)
	col = flag.Int(
		"k",
		0,
		"column index to split",
	)
	header = flag.Bool(
		"h",
		false,
		"if with header",
	)
)

func main() {
	flag.Parse()
	if *in == "" {
		flag.Usage()
		log.Fatal("-i required!")
	}
	if *prefix == "" {
		*prefix = *in
	}
	var fhMap = make(map[string]*os.File)

	var input = textUtil.File2Slice(*in, "\t")
	var title = input[0]

	for i, strings := range input {
		if *header && i == 0 {
			continue
		}
		var key = strings[*col-1]
		var fh, ok = fhMap[key]
		if !ok {
			var output = *prefix + "." + key + ".tsv"
			fh = osUtil.Create(output)
			fmt.Println(output)
			if *header {
				fmtUtil.FprintStringArray(fh, title, "\t")
			}
			fhMap[key] = fh
		}
		fmtUtil.FprintStringArray(fh, strings, "\t")
	}
	for _, file := range fhMap {
		simpleUtil.CheckErr(file.Close())
	}
}
