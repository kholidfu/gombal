package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/mreiferson/go-httpclient"
)

func main() {
	// readFile()
	images := readFile()
	for _, image := range images {
		image = strings.TrimRight(image, "\r\n")
		fmt.Printf("downloading %s\n", image)
		fname := getFileName(image)
		err := download("/tmp/"+fname, image)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func getFileName(url string) string {
	return path.Base(url)
}

var timeout = time.Duration(2 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

func download(fpath, url string) (err error) {
	// Create the file
	out, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer out.Close()

	// set http request with timeout for each request
	transport := &httpclient.Transport{
		ConnectTimeout:        1 * time.Second,
		RequestTimeout:        10 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
	}
	defer transport.Close()
	// Get the data
	client := &http.Client{Transport: transport}
	req, _ := http.NewRequest("GET", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		return err
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
	} else {
		fmt.Println("file not exist")
	}
	return images
}
