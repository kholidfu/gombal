package main

import (
	"fmt"
	"log"
	"net/http"
)

func checkStatus(url string, ch chan<- int) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	ch <- resp.StatusCode
}

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.detik.com",
		"https://www.yahoo.com",
		"https://www.bing.com",
		"https://www.duckduckgo.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.python.org",
	}
	ch := make(chan int)
	for _, url := range urls {
		go checkStatus(url, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
}
