package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func f(n int) int {
	fn := make(map[int]int)
	for i := 0; i <= n; i++ {
		var ft int
		switch i {
		case 0:
			ft = 0
		case 1, 2:
			ft = 1
		default:
			ft = fn[i-1] + fn[i-2]
		}
		fn[i] = ft
	}
	return fn[n]
}

func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			fmt.Println(err)
			os.Exit(100)
		}

		temp := strings.TrimSpace(netData)
		if temp == "STOP" {
			break
		}

		fibo := "-1\n"
		n, err := strconv.Atoi(temp)
		if err == nil {
			fibo = strconv.Itoa(f(n)) + "\n"
		}
		c.Write([]byte(fibo))
	}
	time.Sleep(5 * time.Second)
	c.Close()
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a port number!")
		return
	}

	PORT := ":" + os.Args[1]
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(c)
	}
}
