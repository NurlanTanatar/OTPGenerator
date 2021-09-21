package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func (ip IPAddr) Stringer() {

	for _, val := range ip[:len(ip)-1] {
		fmt.Print(val)
		fmt.Print(".")
	}
	fmt.Println(ip[len(ip)-1])
}

// func main() {
// 	hosts := map[string]IPAddr{
// 		"loopback":  {127, 0, 0, 1},
// 		"googleDNS": {8, 8, 8, 8},
// 	}
// 	for _, ip := range hosts {
// 		//fmt.Printf("%v: %v\n", name, ip)
// 		ip.Stringer()
// 	}
// }
