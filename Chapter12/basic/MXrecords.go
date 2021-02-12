package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Need a domain name!")
		return
	}

	domain := os.Args[1]
	MXs, err := net.LookupMX(domain)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, MX := range MXs {
		fmt.Println(MX.Host)
	}
}
