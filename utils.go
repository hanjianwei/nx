package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
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
