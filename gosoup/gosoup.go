package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/anaskhan96/soup"
)

// This code will create HTTP request to google search result
// and parse the HTML DOM with soup package
// extract title and link
func main() {
	fmt.Println("beautifulsoup alike in Go")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://google.com/search?q=pengacara+yunadi+tersangka", nil)
	if err != nil {
		panic(err.Error())
	}
	// set custom user-agent
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	// do the HTTP request
	resp, err := client.Do(req)
	// fmt.Println(resp)
	if err != nil {
		os.Exit(1)
	}
	// convert back the resp to string
	bytes, err := ioutil.ReadAll(resp.Body)
	doc := soup.HTMLParse(string(bytes))
	// fing and loop through
	links := doc.FindAll("h3", "class", "r")
	for _, link := range links {
		l := link.Find("a")
		if l.Text() != "" {
			fmt.Println(l.Text(), "| Link :", l.Attrs()["href"])
		}
	}
}
