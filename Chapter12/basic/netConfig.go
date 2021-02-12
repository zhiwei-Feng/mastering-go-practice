package main

import (
	"fmt"
	"net"
)

func main() {
	// 获取所有的网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, i := range interfaces {
		fmt.Printf("Interface: %v\n", i.Name)
		// 根据对应接口名获取接口信息
		//byName, err := net.InterfaceByName(i.Name)
		//if err != nil {
		//fmt.Println(err)
		//}
		//addresses, err := byName.Addrs()
		addresses, _ := i.Addrs()
		for k, v := range addresses {
			fmt.Printf("Interface Address #%v: %v\n", k, v.String())
		}
		fmt.Println()
	}
}
