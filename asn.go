package main

import (
	"fmt"
	"log"
	"net"
	"strings"
)

func trimRADBOutput(out string) []string {
	lines := strings.Split(out, "\n")

	if len(lines) > 2 {
		return strings.Split(lines[1], " ")
	}

	return nil
}

// Query ASN information from radb.
// See: http://www.radb.net/support/query2.php
func companyASNs(company string) []string {
	out, err := runCommand(fmt.Sprintf("whois -h whois.radb.net '!iAS-%s,1'", strings.ToUpper(company)))
	if err != nil {
		log.Fatal(err)
	}

	return trimRADBOutput(string(out))
}

func companyIPs(company string) []*net.IPNet {
	if company == "aws" {
		return awsIPs()
	}
	if company == "cf" {
		return cfIPs()
	}

	asns := companyASNs(company)

	var ips []*net.IPNet

	for _, asn := range asns {
		ips = append(ips, asnIPs(asn)...)
	}

	return ips
}

func asnIPs(asn string) []*net.IPNet {
	asn = strings.ToLower(asn)
	if !strings.HasPrefix(asn, "as") {
		asn = "as" + asn
	}

	out, err := runCommand(fmt.Sprintf("whois -h whois.radb.net '!g%s'", asn))
	if err != nil {
		log.Fatal(err)
	}

	return parseIPNets(trimRADBOutput(string(out)))
}
