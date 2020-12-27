package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"regexp"
)

func findIPv4(input string) string {
	partIP := "(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])"
	grammar := partIP + "\\." + partIP + "\\." + partIP + "\\." + partIP
	matchMe := regexp.MustCompile(grammar)
	return matchMe.FindString(input)
}

func main() {
	args := os.Args
	if len(args) < 2 {
		fmt.Printf("usage: %s logFile\n", filepath.Base(args[0]))
		os.Exit(1)
	}

	for _, filename := range args[1:] {
		func() {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Printf("error opening file %s\n", err)
				os.Exit(-1)
			}
			defer f.Close()

			r := bufio.NewReader(f)
			for {
				line, err := r.ReadString('\n')
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Printf("error reading file %s", err)
					break
				}

				ip := findIPv4(line)
				trial := net.ParseIP(ip)
				if trial.To4() == nil {
					continue
				} else {
					fmt.Println(ip)
				}
			}
		}()
	}
}
