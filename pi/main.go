/*
Find the location of a six-digit number at PI.
Usage: $ ./pi 123456
*/
package main

import (
	"fmt"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("/Users/peter/Work/pi-billion.txt")
	var readBytes int64
	var cycles int64
	var i int64
	var searchStr string

	if len(os.Args) > 1 {
		searchStr = os.Args[1]
	} else {
		searchStr = "000000"
	}

	readBytes = 1000
	cycles = 1000000
	bytes := make([]byte, readBytes)

	for i = 0; i <= cycles; i++ {
		if i == 0 {
			_, err = file.Seek(0, 0)
		} else {
			_, err = file.Seek(i*(readBytes-5), 0)
		}
		check(err)
		_, err = file.Read(bytes)
		check(err)

		str := string(bytes)
		index := strings.Index(str, searchStr)
		if index >= 0 {
			location := i*(readBytes-5) + int64(index) - 1
			fmt.Printf("\"%s\" is located at %d \n%s\n", searchStr,
				location, string(bytes))
			break
		}
	}
}
