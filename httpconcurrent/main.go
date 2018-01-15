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
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/gosimple/slug"
	"github.com/mreiferson/go-httpclient"
)

func main() {
	images := readFile()
	for _, image := range images {
		// create downloaded dir if not exist
		downloadPath := getWorkingDir() + "downloaded"
		if _, err := os.Stat(downloadPath); err != nil {
			// if not exist, mkdir
			err := os.Mkdir(downloadPath, 0777)
			if err != nil {
				fmt.Println(err)
			}
		}
		image = strings.TrimRight(image, "\r\n")
		fmt.Printf("downloading %s\n", image)
		fname := getFileName(image)
		err := download(getWorkingDir()+"downloaded"+getSeparator()+fname, image)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}

func checkDirExist(fpath string) bool {
	dirExist := true
	if _, err := os.Stat(fpath); err != nil {
		fmt.Println("dir not exist")
		dirExist = false
	}
	return dirExist
}

func getSeparator() string {
	var sep string
	myOS := runtime.GOOS

	if myOS == "windows" {
		sep = "\\"
	} else {
		sep = "/"
	}
	return sep
}

func getFileName(url string) string {
	rawFName := path.Base(url)
	ext := filepath.Ext(rawFName)
	fname := strings.Split(rawFName, ext)
	sluggedFName := slug.Make(fname[0])
	return sluggedFName + ext
}

func getWorkingDir() string {
	var wd string
	myOS := runtime.GOOS
	targetDir, _ := os.Getwd()

	if myOS == "windows" {
		wd = strings.Join([]string{targetDir, "src", "httpconcurrent" + getSeparator()}, "\\")
	} else {
		wd = strings.Join([]string{targetDir, "src", "httpconcurrent" + getSeparator()}, "/")
	}
	return wd
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

	wd := getWorkingDir()
	pathToFile := wd + "images.txt"

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
