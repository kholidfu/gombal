package main

import (
	"fmt"
)

// Kitchen struct
type Kitchen struct {
	numOfForks  int
	numOfKnives int
}

// House struct
// House is embedded in Kitchen, so House have access to Kitchen property
type House struct {
	Kitchen
}

// func total
// we can access this func through House struct
// create House object then call the total method
func (ht House) total() int {
	totalTool := ht.numOfForks + ht.numOfKnives
	return totalTool
}

func main() {
	h := House{Kitchen{4, 4}}
	fmt.Println(h.numOfForks)
	fmt.Println(h.total())
}
