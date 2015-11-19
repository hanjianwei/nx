package nx

// Resolve given name
func Resolve(name string) {
	if isDigits(name) {
		printIPNets(asnIPs(name))
	} else if isDomain(name) {
		printIPs(domainIPs(name))
	} else {
		printIPNets(companyIPs(name))
	}
}
