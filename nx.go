package main

import (
	"flag"
	"fmt"
)

const appVersion = "0.1.0"

func resolve(name string) {
	if isDigits(name) {
		printIPNets(asnIPs(name))
	} else if isDomain(name) {
		printIPs(domainIPs(name))
	} else {
		printIPNets(companyIPs(name))
	}
}

func main() {
	version := flag.Bool("v", false, "print version")

	flag.Parse()

	if *version {
		fmt.Println(appVersion)
	} else {
		for _, arg := range flag.Args() {
			resolve(arg)
		}
	}
}
