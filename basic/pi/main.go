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
	const ReadBytes = 256
	const ReadCycles = 1000000000 / 256

	var searchStr string
	bytes := make([]byte, ReadBytes)

	if len(os.Args) > 1 {
		searchStr = os.Args[1]
	} else {
		searchStr = "000000"
	}
	file, err := os.Open("/Users/peter/Work/pi-billion.txt")
	check(err)
	defer file.Close()

	for i := 0; i <= ReadCycles; i++ {
		if i == 0 {
			_, err = file.Seek(0, 0)
		} else {
			_, err = file.Seek(int64(i*(ReadBytes-5)), 0)
		}
		check(err)
		_, err = file.Read(bytes)
		check(err)

		str := string(bytes)
		index := strings.Index(str, searchStr)
		if index >= 0 {
			location := i*(ReadBytes-5) + index - 1
			fmt.Printf("\"%s\" is located at %d \n%s\n", searchStr,
				location, string(bytes))
			break
		}
	}
}
