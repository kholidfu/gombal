package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func downloadFile(fpath, url string) (err error) {
	// Create the file
	out, err := os.Create(fpath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
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

func main() {
	err := downloadFile("./src/downloader/gopher.jpg", "https://vorozhko.net/wp-content/uploads/2017/11/gophercon2015.jpg")
	if err != nil {
		log.Fatal(err)
	}
}
