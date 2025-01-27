package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func countWords(url string, frequency map[string]int) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("Server returning error status code: " + resp.Status)
	}
	body, _ := io.ReadAll(resp.Body)
	wordRegex := regexp.MustCompile(`[a-zA-Z]+`)
	for _, word := range wordRegex.FindAllString(string(body), -1) {
		wordLower := strings.ToLower(string(word))
		frequency[wordLower] += 1
	}
	fmt.Println("Completed:", url)
}

func main() {
	frequency := make(map[string]int)
	for i := 1000; i <= 1020; i++ {
		url := fmt.Sprintf("https://rfc-editor.org/rfc/rfc%d.txt", i)
		go countWords(url, frequency)
	}
	time.Sleep(10 * time.Second)
	for k, v := range frequency {
		fmt.Println(k, "->", v)
	}
}
