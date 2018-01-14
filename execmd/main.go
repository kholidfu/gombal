package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	myOS := runtime.GOOS
	var err error
	var cmd []byte
	if myOS == "windows" {
		cmd, err = exec.Command("cmd", "/C", "dir").Output()
	} else {
		cmd, err = exec.Command("ls").Output()
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(cmd))
}
