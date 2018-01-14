package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/anaskhan96/soup"
	"github.com/tidwall/gjson"
)

func main() {
	url := "https://www.google.com/search?hl=en&site=imghp&tbm=isch&tbs=isz:l&q=golang+wallpaper"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	doc := soup.HTMLParse(string(bytes))
	// fing and loop through
	data := doc.FindAll("div", "class", "rg_meta notranslate")
	for _, d := range data {
		u := gjson.Get(d.Text(), "ou")
		fmt.Println(u)
	}
}
