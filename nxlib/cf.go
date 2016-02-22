package nx

import (
	"log"
	"net"
	"strings"
)

func cfIPs() []*net.IPNet {
	url := "https://www.cloudflare.com/ips-v4"
	content, err := readURL(url)
	if err != nil {
		log.Fatal(err)
	}

	return parseIPNets(strings.Fields(string(content)))
}
