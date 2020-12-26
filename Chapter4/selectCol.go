package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: SelectColumn <column> <file1> [<file2> [...<fileN]]\n")
		os.Exit(1)
	}
	// args []
	temp, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Column value is not an integer:", temp)
		return
	}

	column := temp
	if column < 0 {
		fmt.Println("Invalid Column number!")
		os.Exit(1)
	}

	for _, filename := range args[2:] {
		fmt.Println("\t\t", filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Printf("error opening file %s\n", err)
			continue
		}
		defer f.Close()
		r := bufio.NewReader(f)
		for {
			// 这种方法会出现一个问题，即无法处理 不是以换行符结尾的最后一行字符串
			// 见https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
			line, err := r.ReadString('\n')
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Printf("error reading file %s", err)
			}
			data := strings.Fields(line)
			if len(data) >= column {
				fmt.Println(data[column-1])
			}
		}
	}
}
