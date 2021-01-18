package main

import (
	"Chapter6/a"
	"Chapter6/b"
	"fmt"
)

func init() {
	fmt.Println("init() manyInit")
}

func main() {
	a.FromA()
	b.FromB()
}
