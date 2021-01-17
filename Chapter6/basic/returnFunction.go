package main

import "fmt"

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}

func main() {
	j := funReturnFun()
	i := funReturnFun()

	fmt.Println("1:", i())
	fmt.Println("2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
	fmt.Println("3:", i())
}
