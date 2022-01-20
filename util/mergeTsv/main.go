package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/liserjrqlxue/goUtil/textUtil"
)

var (
	lst = flag.String(
		"l",
		"",
		"input list",
	)
	title = flag.String(
		"t",
		"",
		"output title from file, default is first file",
	)
)

func main() {
	flag.Parse()
	var inputList []string
	if *lst != "" {
		inputList = textUtil.File2Array(*lst)
	}
	inputList = append(inputList, flag.Args()...)
	if *title == "" {
		*title = inputList[0]
	}
	var _, titles = textUtil.File2MapArray(*title, "\t", nil)
	fmt.Println(strings.Join(titles, "\t"))
	for _, s := range inputList {
		var ma, _ = textUtil.File2MapArray(s, "\t", nil)
		for _, m := range ma {
			var output []string
			for _, t := range titles {
				output = append(output, m[t])

			}
			fmt.Println(strings.Join(output, "\t"))
		}
	}
}
