package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// stationname struct
type stationname struct {
	// ID          int    `json:"id"`
	StationName string `json:"stationName"`
}

// citibank struct
type citibank struct {
	ExecutionTime      string        `json:"executiontime"`
	ArrStationBeanList []stationname `json:"stationBeanList"`
}

var url string

func main() {
	url = "https://feeds.citibikenyc.com/stations/stations.json"

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
		bank := citibank{}
		err := json.Unmarshal(b, &bank)
		if err != nil {
			panic(err.Error())
		}
		// fmt.Println(ip1.ArrStationBeanList[0])
		addresses := bank.ArrStationBeanList
		for _, v := range addresses {
			fmt.Println(v.StationName)
		}
	}
}
