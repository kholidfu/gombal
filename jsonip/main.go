package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IP struct
type ip struct {
	IP    string `json:"ip"`
	Quote string `json:"reject-fascism"`
}

func main() {
	url := "https://jsonip.com/"
	// open HTTP request to url
	r, err := http.Get(url)
	// check error
	if err != nil {
		panic(err.Error())
	}
	defer r.Body.Close()
	// read the body
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	// if HTTP status 200
	if r.StatusCode == 200 {
		// print the response body
		// fmt.Println(string(b))
		ip1 := ip{}
		err := json.Unmarshal(b, &ip1)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(ip1.IP)
		fmt.Println(ip1.Quote)
	}
}
