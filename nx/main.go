package main

import (
	"flag"
	"fmt"

	"github.com/hanjianwei/nx"
)

const appVersion = "0.1.0"

func main() {
	version := flag.Bool("v", false, "print version")

	flag.Parse()

	if *version {
		fmt.Println(appVersion)
	} else {
		for _, arg := range flag.Args() {
			nx.Resolve(arg)
		}
	}
}
