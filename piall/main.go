/*
Find all 6-digit numbers' location at PI
A file named "result" generated
Run Time = 1254.18sec on my Mac
*/
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const READ_BYTES int64 = 1024 * 4
	const READ_CYCLES int64 = 1000000000/READ_BYTES + 1
	var i int64 = 0
	start := time.Now()

	file, err := os.Open("/Users/peter/Work/pi-billion.txt")
	check(err)
	defer file.Close()

	// to search all of the 000000~999999 numbers
	numbers := make(map[string]bool)
	for k := 0; k < 1e6; k++ {
		key := fmt.Sprintf("%06d", k)
		numbers[key] = false
	}

	// result write to
	result, err := os.Create("result")
	check(err)
	defer result.Close()

	bytes := make([]byte, READ_BYTES)
	for i = 0; i <= READ_CYCLES; i++ {
		if i == 0 {
			_, err = file.Seek(0, 0)
		} else {
			_, err = file.Seek(i*(READ_BYTES-5), 0)
		}
		check(err)
		_, err = file.Read(bytes)
		check(err)
		str := string(bytes)

		for searchNum := range numbers {
			index := strings.Index(str, searchNum)
			if index >= 0 {
				result.WriteString(fmt.Sprintf("\"%s\" at %d\n", searchNum, i*(READ_BYTES-5)+int64(index)-1))
				delete(numbers, searchNum)
			}
		}
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("Done.  map size = %d.  Time = %.2fsec\n", len(numbers), secs)
}
