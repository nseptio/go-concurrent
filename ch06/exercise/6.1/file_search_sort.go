package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	wg.Add(1)
	listFiles := make([]string, 0)
	go fileSearch(os.Args[1], os.Args[2], &wg, &mutex, &listFiles)
	wg.Wait()

	mutex.Lock()
	slices.SortFunc(listFiles, func(a, b string) int {
		return strings.Compare(b, a)
	})
	fmt.Println(strings.Join(listFiles, "\n"))
	mutex.Unlock()
}

func fileSearch(dir, filename string, wg *sync.WaitGroup, mutex *sync.Mutex, listFiles *[]string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(files); i++ {
		fpath := filepath.Join(dir, files[i].Name())
		if strings.Contains(fpath, filename) {
			mutex.Lock()
			*listFiles = append(*listFiles, fpath)
			mutex.Unlock()
		}

		if files[i].IsDir() {
			wg.Add(1)
			go fileSearch(fpath, filename, wg, mutex, listFiles)
		}
	}

	wg.Done()
}
