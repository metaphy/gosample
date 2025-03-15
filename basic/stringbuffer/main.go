/*
StringBuffer
*/
package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.WriteByte('[') // Write a byte
	for i := 0; i < 10; i++ {
		// Write string
		buf.WriteString(fmt.Sprintf("%06d", i))
		if i != 9 {
			buf.WriteByte(',')
		}
	}

	// fmt write to buf
	fmt.Fprintf(&buf, "]")
	fmt.Println(buf.String())
}
