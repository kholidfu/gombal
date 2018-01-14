package main

import (
	"log"
	"os"
	"runtime"
	"strings"
)

// CreateFile func
func CreateFile(fname string) {
	// detect os
	myOS := runtime.GOOS
	targetDir, _ := os.Getwd()

	// DECLARE var before if statement
	var pathToFile string

	if myOS == "windows" {
		// ASSIGN here
		pathToFile = strings.Join([]string{targetDir, "src", "fileops", fname}, "\\")
	}

	// or ASSIGN here
	pathToFile = strings.Join([]string{targetDir, "src", "fileops", fname}, "/")
	newFile, err := os.Create(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(newFile)
}

func main() {
	CreateFile("go.txt")
}
