package nx

import (
	"encoding/json"
	"log"
	"net"
	"strings"
)

func awsIPs() []*net.IPNet {
	// Fetch IP ranges
	url := "https://ip-ranges.amazonaws.com/ip-ranges.json"
	content, err := readURL(url)
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON
	var res struct {
		SyncToken  string `json:"syncToken"`
		CreateDate string `json:"createDate"`
		Prefixes   []struct {
			IPPrefix string `json:"ip_prefix"`
			Region   string `json:"region"`
			Service  string `json:"service"`
		} `json:"prefixes"`
	}
	if err := json.Unmarshal(content, &res); err != nil {
		log.Fatal(err)
	}

	// Extract IP ranges
	var ipstrs []string
	for _, prefix := range res.Prefixes {
		if !strings.HasPrefix(prefix.Region, "cn-") {
			ipstrs = append(ipstrs, prefix.IPPrefix)
		}
	}

	return parseIPNets(ipstrs)
}
