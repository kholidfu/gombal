package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	readFile()
	images := readFile()
	for _, image := range images {
		image = strings.TrimRight(image, "\r\n")
		fmt.Printf("downloading %s\n", image)
		fname := getFileName(image)
		err := download("/tmp/"+fname, image)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getFileName(url string) string {
	return path.Base(url)
}

func download(fpath, url string) (err error) {
	// Create the file
	out, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer out.Close()

	// set timeout
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	// Get the data
	resp, err := client.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Writer the response body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// read from file
// return list of url to the image
func readFile() []string {

	var images []string
	var line string

	pathToFile := strings.Join([]string{"/home", "kholidfu", "Desktop", "images.txt"}, "/")
	// check if file exist
	if _, err := os.Stat(pathToFile); err == nil {
		// read the file
		f, err := os.Open(pathToFile)
		if err != nil {
			log.Fatal(err)
		}
		reader := bufio.NewReader(f)
		for {
			line, err = reader.ReadString('\n')
			// check if line not blank
			if len(line) > 0 {
				images = append(images, line)
			}
			if err != nil {
				break
			}
		}
	}
	return images
}
