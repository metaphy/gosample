/*
Find all 6-digit numbers' location at PI
A file named "result" generated
Run Time = 1358.19sec on my Mac
*/
package main

import (
	"bytes"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const ReadBytes int = 1024 * 4
	const ReadCycles int = 1000000000/ReadBytes + 1
	var cycle int
	var buf bytes.Buffer

	start := time.Now()
	fmt.Println(start)
	// to search all of the 000000~999999 numbers
	numbers := make(map[int]bool)
	for k := 0; k < 1e6; k++ {
		numbers[k] = false
	}

	// Open the file PI
	file, err := os.Open("/Users/peter/Work/pi-billion.txt")
	check(err)
	defer file.Close()

	// result write to
	result, err := os.Create("result-sorted.txt")
	check(err)
	defer result.Close()

	bytes := make([]byte, ReadBytes)
	for cycle = 0; cycle <= ReadCycles; cycle++ {
		if cycle == 0 {
			_, err = file.Seek(0, 0)
		} else {
			_, err = file.Seek(int64(cycle*(ReadBytes-5)), 0)
		}
		check(err)
		_, err = file.Read(bytes)
		check(err)

		var foundLocs []int                  // sorted locations
		mapFoundNumbers := make(map[int]int) // map[location]number

		str := string(bytes)
		for num := range numbers {
			numStr := fmt.Sprintf("%06d", num)
			index := strings.Index(str, numStr)
			if index >= 0 {
				loc := cycle*(ReadBytes-5) + index - 1
				foundLocs = append(foundLocs, loc)
				mapFoundNumbers[loc] = num
				delete(numbers, num)
			}
		}
		sort.Ints(foundLocs)

		for _, loc := range foundLocs {
			numStr := fmt.Sprintf("%06d", mapFoundNumbers[loc])
			buf.WriteString(fmt.Sprintf("\"%s\" : %d\n", numStr, loc))
		}
	}
	result.WriteString(buf.String())
	secs := time.Since(start).Seconds()
	fmt.Printf("Done.  map size = %d.  Time = %.2fsec\n", len(numbers), secs)
}
