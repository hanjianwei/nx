package nx

import (
	"log"
	"net"

	"github.com/miekg/dns"
)

func domainIPs(domain string) []net.IP {
	m := new(dns.Msg)
	m.SetQuestion(dns.Fqdn(domain), dns.TypeA)

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
