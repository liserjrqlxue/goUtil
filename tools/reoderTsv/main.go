package main

import (
	"bufio"
	"flag"
	"strings"

	"github.com/liserjrqlxue/goUtil/fmtUtil"
	"github.com/liserjrqlxue/goUtil/osUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

var (
	input = flag.String(
		"input",
		"",
		"input tsv, required",
	)
	output = flag.String(
		"output",
		"",
		"output tsv, default is -input.reorder",
	)
	first = flag.String(
		"first",
		"",
		"columns move to first, sep by comma",
	)
	order = flag.String(
		"order",
		"",
		"order columns, sep by comma",
	)
	sep = flag.String(
		"sep",
		"\t",
		"sep of input",
	)
)

func main() {
	flag.Parse()
	if *input == "" {
		flag.Usage()
		return
	}
	if *output == "" {
		*output = *input + ".reorder"
	}
	var inFirst = make(map[string]bool)
	var titles []string
	if *first != "" {
		for _, k := range strings.Split(*first, ",") {
			titles = append(titles, k)
			inFirst[k] = true
		}
	}
	if *order != "" {
		for _, k := range strings.Split(*order, ",") {
			if !inFirst[k] {
				titles = append(titles, k)
			}
		}
	}

	var inputF = osUtil.Open(*input)
	defer simpleUtil.DeferClose(inputF)

	var outputF = osUtil.Create(*output)
	defer simpleUtil.DeferClose(outputF)

	var scanner = bufio.NewScanner(inputF)

	var header = true
	var columns []string
	for scanner.Scan() {
		var line = scanner.Text()
		var a = strings.Split(line, *sep)
		if header {
			header = false
			columns = a
			// set first line as default -order
			if *order == "" {
				for _, k := range a {
					if !inFirst[k] {
						titles = append(titles, k)
					}
				}
			}
			fmtUtil.FprintStringArray(outputF, titles, *sep)
		} else {
			var item = make(map[string]string)
			for i, k := range a {
				item[columns[i]] = k
			}
			var outputArray []string
			for _, k := range titles {
				outputArray = append(outputArray, item[k])
			}
			fmtUtil.FprintStringArray(outputF, outputArray, *sep)
		}
	}
	simpleUtil.CheckErr(scanner.Err())
}
