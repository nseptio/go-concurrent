// Adapt the program in the third exercise to continue searching recursively in
// any subdirectories. If you give your search goroutine a file, it should search for a
// string match in that file, just like in the previous exercises. Otherwise, if you
// give it a directory, it should recursively spawn a new goroutine for each file or
// directory found inside. Call the program grepdirrec.go, and execute it by run-
// ning this command:
// go run grepdirrec.go bubbles ../../commonfiles

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		panic("usage go run grepdirrec.go <word> <dir path>")
	}

	searchWord := os.Args[1]
	dirPath := os.Args[2]

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	for _, entry := range entries {
		go grepRec(dirPath, entry, searchWord)
	}

	time.Sleep(3 * time.Second)
}

func grepRec(path string, entry os.DirEntry, searchWord string) {
	fullPath := filepath.Join(path, entry.Name())
	if entry.IsDir() {
		files, err := os.ReadDir(fullPath)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			go grepRec(fullPath, file, searchWord)
		}
	} else {
		file, err := os.Open(fullPath)
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
			fmt.Printf("%s contains word %s\n", fullPath, searchWord)
		} else {
			fmt.Printf("%s does NOT contains %s\n", fullPath, searchWord)
		}
	}
}
