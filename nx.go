package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

const appVersion = "0.1.0"

func isDigits(name string) bool {
	if m, _ := regexp.MatchString(`^[1-9][0-9]*$`, name); !m {
		return false
	}
	return true
}

func isDomain(name string) bool {
	return strings.Contains(name, ".")
}

func resolve(name string) {
	if isDigits(name) {
		fmt.Println(asnIPs(name))
	} else if isDomain(name) {
		fmt.Println(domainIPs(name))
	} else {
		fmt.Println(companyIPs(name))
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
