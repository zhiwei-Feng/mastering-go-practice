package main

import (
	sharedRPC "Chapter13/rpc"
	"fmt"
	"net/rpc"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a host:port string!")
		return
	}

	connect := os.Args[1]
	c, err := rpc.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return
	}

	args := sharedRPC.MyFloats{16, -0.5}
	var reply float64

	err = c.Call("MyInterface.Multiply", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("MyInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
