package main

import (
	"flag"
	"fmt"
)

const appVersion = "0.1.0"

func main() {
	version := flag.Bool("v", false, "print version")

	flag.Parse()

	if *version {
		fmt.Println(appVersion)
	}

	fmt.Println(flag.Args())
}
