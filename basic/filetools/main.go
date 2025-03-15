package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// write to a file
	file, err := os.Create("test.txt")
	check(err)
	for i := 0; i < 5; i++ {
		file.WriteString(fmt.Sprintf("hello %06d\n", i))
	}
	file.Close()

	// Read the file to bytes
	data, err := ioutil.ReadFile("test.txt")
	check(err)
	fmt.Println("data:", data)
	fmt.Print("String of data:\n", string(data))

	// Read lines of the file
	file, err = os.Open("test.txt")
	check(err)
	input := bufio.NewScanner(file)
	fmt.Println("Read lines of the file:")
	for input.Scan() {
		fmt.Println(input.Text())
	}
	file.Close()

	// Reader. read to bytes
	file, err = os.Open("test.txt")
	reader := bufio.NewReader(file)
	bytes, err := reader.Peek(10)
	check(err)
	fmt.Printf("The first 10 bytes of the file: %q\n", bytes)
	file.Close()

	// Read file from somewhere
	file, err = os.Open("test.txt")
	_, err = file.Seek(10, 0)
	check(err)
	bytes = make([]byte, 20)
	readnum, err := file.Read(bytes)
	check(err)
	fmt.Printf("The %d bytes from the 10th: %s\n", readnum, string(bytes))
}
