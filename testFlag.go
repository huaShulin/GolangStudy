package main

import (
	"flag"
	"fmt"
)

var n = flag.Bool("n", false, "")

//var n = flag.Bool("n", true, "")
var s = flag.String("s", "Hi", "")

func main() {
	//很重要
	flag.Parse()

	if *n {
		fmt.Println(*s)
	}
}
