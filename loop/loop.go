package main

import (
	"fmt"
)

func main() {
	// array/list example
	l := []int{1, 2, 3, 4, 5}

	for _, i := range l {
		fmt.Println(i)
	}

	// map example (dict in python)
	d := map[string]int{
		"kholidfu": 38,
		"mamah":    36,
	}

	for k, v := range d {
		fmt.Println(k)
		fmt.Println(v)
	}
}
