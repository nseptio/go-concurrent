// Change the program you wrote in the second exercise so that instead of passing
// a list of text filenames, you pass a directory path. The program will look inside
// this directory and list the files. For each file, you can spawn a goroutine that will
// search for a string match (the same as before). Call the program grepdir.go.
// Here’s how you can execute this Go program:
// go run grepdir.go bubbles ../../commonfiles

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		panic("usage go run grepdir.go <word> <dir>")
	}

	searchWord := os.Args[1]
	dirPath := os.Args[2]

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			fmt.Printf("Skipping directory: %s\n", entry.Name())
			continue
		}

		filePath := dirPath + "/" + entry.Name()
		go grepDir(filePath, searchWord)

		fmt.Println()
	}

	time.Sleep(3 * time.Second)
}

func grepDir(filename, searchWord string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	isContain := false
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchWord) {
			isContain = true
		}
	}

	if isContain {
		fmt.Printf("%s contains word %s\n", filename, searchWord)
	} else {
		fmt.Println(false)
	}
}
