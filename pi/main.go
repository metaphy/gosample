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

func index(searchStr string, bytes []byte) int {
	for i := 0; i <= len(bytes)-len(searchStr); i++ {
		if searchStr[0] == bytes[i] && searchStr[1] == bytes[i+1] &&
			searchStr[2] == bytes[i+2] && searchStr[3] == bytes[i+3] &&
			searchStr[4] == bytes[i+4] && searchStr[5] == bytes[i+5] {
			return i
		}
	}
	return -1
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
		if strings.Contains(str, searchStr) {
			relative := index(searchStr, bytes)
			fmt.Printf("\"%s\" is located at %d \n%s\n", searchStr,
				i*(readBytes-5)+int64(relative)-1, string(bytes))
			break
		}
	}
}
