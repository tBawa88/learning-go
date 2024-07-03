package main

import (
	"flag"
	"fmt"
	"strings"
)

var sep = flag.String("O", ",", "seperator for command line arguments") //returns *string, default value is ","
var n = flag.Bool("no", false, "omit trailing newline")                 //returns *bool, default value is false

func echo() {
	flag.Parse()

	cla := strings.Join(flag.Args(), *sep)
	fmt.Println(cla)
	if !*n {
		fmt.Println()
	}
}
