package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go fileSearch(os.Args[1], os.Args[2], &wg)
	wg.Wait()
}

func fileSearch(dir, filename string, wg *sync.WaitGroup) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(files); i++ {
		fpath := filepath.Join(dir, files[i].Name())
		if strings.Contains(fpath, filename) {
			fmt.Println(fpath)
		}

		if files[i].IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg)
		}
	}

	wg.Done()
}
