package main

import (
	"fmt"
	"strings"
)

//Describe the differences between a character, a byte, and a rune
func main() {
	// reference: https://blog.golang.org/strings
	ansBuilder := strings.Builder{}
	ansBuilder.WriteString("个人理解如下：\n")
	ansBuilder.WriteString("byte -> 8位二进制数据, e.g. \\xbd\n")
	ansBuilder.WriteString("rune -> unicode code point, e.g. U+2318\n")
	ansBuilder.WriteString("character -> 一个由不同code point组成的数, e.g. à由U+0300和U+0061组成\n")
	fmt.Println(ansBuilder.String())
}
