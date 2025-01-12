package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 3 {
		panic("Usage: go run grepfiles.go <word> <file1> ...")
	}

	fileNames := os.Args[2:]
	searchWord := os.Args[1]
	fmt.Println(fileNames, searchWord)

	for _, fileName := range fileNames {
		go grepFile(fileName, searchWord)
	}

	time.Sleep(3 * time.Second)
}

func grepFile(fileName, searchWord string) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	isContains := false

	for {
		line, _, err := reader.ReadLine()
		if len(line) > 0 {
			if strings.Contains(string(line), searchWord) {
				isContains = true
				break
			}
		}

		if err != nil {
			break
		}
	}

	fmt.Println(isContains)
}
