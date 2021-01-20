package main

import (
	"fmt"
	"os"
	"syscall"
)

// 这个程序在go 1.15.2版本下有问题，在windows下运行也有问题
func main() {
	pid := syscall.Getpid()
	fmt.Println("My pid is", pid)
	uid := syscall.Getuid()
	fmt.Println("User ID:", uid)
	message := []byte{'H', 'e', 'l', 'l', 'o', '!', '\n'}
	fd := 1
	syscall.Write(syscall.Handle(fd), message)
	fmt.Println("Using syscall.Exec()")
	command := "/bin/ls"
	env := os.Environ()
	syscall.Exec(command, []string{"ls", "-a", "-x"}, env)
}
