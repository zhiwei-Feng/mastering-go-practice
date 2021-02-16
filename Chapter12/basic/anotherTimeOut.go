package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var timeout = time.Second

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please provide a URL")
		return
	}

	if len(os.Args) == 3 {
		// args[1]: URL, args[2]: timeout
		temp, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Using Default Timeout!")
		} else {
			timeout = time.Duration(temp) * time.Second
		}
	}

	URL := os.Args[1]
	client := http.Client{Timeout: timeout}
	data, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		defer data.Body.Close()
		_, err := io.Copy(os.Stdout, data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
