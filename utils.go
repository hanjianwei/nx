package nx

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os/exec"
	"regexp"
	"strings"
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

func isDigits(name string) bool {
	if m, _ := regexp.MatchString(`^[1-9][0-9]*$`, name); !m {
		return false
	}
	return true
}

func isDomain(name string) bool {
	return strings.Contains(name, ".")
}

func printIPNets(ipnets []*net.IPNet) {
	for _, ipnet := range ipnets {
		fmt.Println(ipnet)
	}
}

func printIPs(ips []net.IP) {
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
