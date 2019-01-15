/*
Find all 6-digit numbers' location at PI
A file named "result" generated
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
	const READ_BYTES int64 = 1000
	const READ_CYCLES int64 = 1000000
	var i int64 = 0
	start := time.Now()

	file, err := os.Open("/Users/peter/Work/pi-billion.txt")
	// to search all of the 000000~999999 numbers
	mp := make(map[string]int64)
	for k := 0; k < 1000000; k++ {
		key := fmt.Sprintf("%06d", k)
		mp[key] = -1
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

		for key := range mp {
			if strings.Contains(str, key) {
				relative := index(key, bytes)
				mp[key] = i*(READ_BYTES-5) + int64(relative) - 1
				result.WriteString(fmt.Sprintf("\"%s\" at %d\n", key, mp[key]))
				delete(mp, key)
			}
		}
	}
	secs := time.Since(start).Seconds()
	fmt.Printf("Done.  map size = %d.  Time = %.2fsec\n", len(mp), secs)
}
