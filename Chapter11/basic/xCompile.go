package main

import (
	"fmt"
	"runtime"
)

func main() {
	// env GOOS=linux GOARCH=386 go build xCompile.go
	fmt.Print("You are using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH, "machine")
	fmt.Println("with Go version", runtime.Version())
}
