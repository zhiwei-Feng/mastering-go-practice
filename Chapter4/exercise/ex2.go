package main

import (
	"fmt"
	"strings"
)

//Describe the differences between make and new without looking at the chapter
func main() {
	ansBuilder := strings.Builder{}
	ansBuilder.WriteString("The main difference between new and make is that\n")
	ansBuilder.WriteString("variables created with make are properly initialized without just zeroing the allocated memory space.\n")
	ansBuilder.WriteString("Additionally, make can only be applied to maps, channels, and slices, and it does not return a memory address\n")
	ansBuilder.WriteString("which means that make does not return a pointer.\n")
	fmt.Println(ansBuilder.String())
}
