package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"

	"github.com/miekg/dns"
)

func readURL(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	content, err := ioutil.ReadAll(res.Body)
	return content, err
}

func runCommand(cmdStr string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	return cmd.Output()
}

func parseIPNets(ipstrs []string) []*net.IPNet {
	ipnets := make([]*net.IPNet, len(ipstrs))
	for i, s := range ipstrs {
		_, ipnet, err := net.ParseCIDR(s)
		if err != nil {
			log.Fatal(err)
		}
		ipnets[i] = ipnet
	}

	return ipnets
}

func resolveDomain(domain string) []net.IP {
	m := new(dns.Msg)
	m.SetQuestion(domain, dns.TypeA)

	in, err := dns.Exchange(m, "8.8.8.8:53")
	if err != nil {
		log.Fatal(err)
	}

	ips := make([]net.IP, len(in.Answer))
	for i, a := range in.Answer {
		ips[i] = a.(*dns.A).A
	}

	return ips
}
