package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// Write a program similar to the one in listing 2.3 that accepts a list of text file-
// names as arguments. For each filename, the program should spawn a new
// goroutine that will output the contents of that file to the console. You can use
// the time.Sleep() function to wait for the child goroutines to complete (until
// you know how to do this better). Call the program catfiles.go. Here’s how you
// can execute this Go program:
// go run catfiles.go txtfile1 txtfile2 txtfile3

func main() {
	if len(os.Args) < 2 {
		panic("Usage: go run catfiles.go <file1> <file2> ...")
	}

	fileNames := os.Args[1:]
	for _, fileName := range fileNames {
		go catFile(fileName)
	}

	time.Sleep(3 * time.Second)
}

func catFile(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			fmt.Println(string(line))
		}

		if err != nil {
			break
		}
	}
}
