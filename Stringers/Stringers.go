// https://go.dev/tour/methods/18


package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ipaddr IPAddr) String() string {
	ip_string:=""
	const dot="."
	for i,ip_digit := range ipaddr {
		ip_string += fmt.Sprint(ip_digit)
		if i==len(ipaddr)-1 {
			break
		}
		ip_string+=dot
	}
	return ip_string
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v ip address is %v \n", name, ip)
	}
}
