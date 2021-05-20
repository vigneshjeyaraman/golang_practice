package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.linkedin.com",
		"https://www.amazon.in",
	}
	ch := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			ch <- resp.Header.Get("Content-Type")
			defer resp.Body.Close()
		}(url)
	}
	i := 0
	for i < len(urls) {
		fmt.Println(<-ch)
		i++
	}
}
