package b

import (
	"Chapter6/a"
	"fmt"
)

func init() {
	fmt.Println("init() b")
}

func FromB() {
	fmt.Println("fromB()")
	a.FromA()
}
