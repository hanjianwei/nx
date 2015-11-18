package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func trimRADBOutput(out string) []string {
	lines := strings.Split(out, "\n")

	if len(lines) > 1 {
		return strings.Split(lines[1], " ")
	}

	return nil
}

// Query ASN information from radb.
// See: http://www.radb.net/support/query2.php
func companyIPs(company string) []string {
	out, err := runCommand(fmt.Sprintf("whois -h whois.radb.net '!i%s,1'", company))
	if err != nil {
		log.Fatal(err)
	}

	return trimRADBOutput(string(out))
}

func asnIPs(asn string) []*net.IPNet {
	out, err := runCommand(fmt.Sprintf("whois -h whois.radb.net '!g%s'", asn))
	if err != nil {
		log.Fatal(err)
	}

	return parseIPNets(trimRADBOutput(string(out)))
}
