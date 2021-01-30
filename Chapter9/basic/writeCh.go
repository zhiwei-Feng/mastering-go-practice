package main

import (
	"fmt"
	"time"
)

func writeToChannel(c chan int, x int) {
	fmt.Println(x)
	c <- x
	close(c)
	// 实际上这个goroutine并没有结束，因为没有channel的读操作,导致block
	fmt.Println(x)
}

func main() {
	c := make(chan int)
	go writeToChannel(c, 10)

	time.Sleep(1 * time.Second)
}
