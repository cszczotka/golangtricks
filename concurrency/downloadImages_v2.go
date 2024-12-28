package main

import (
 "fmt"
 "io"
 "mime"
 "net/http"
 "os"
 "path/filepath"
 "sync"
)

type DownloadResult struct {
	URL   string
	Error error
}

func getFileNameFromHeader(header http.Header) (string, error) {
	contentDisposition := header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
	 return "", err
	}
	return params["filename"], nil
   }
   
   func getFileNameFromURL(url string) string {
	// Use the last segment of the URL as the filename
	_, fileName := filepath.Split(url)
	return fileName
   }
   
   func worker(id int, wg *sync.WaitGroup, input <-chan string, results chan<- DownloadResult) {
	defer wg.Done()
   
	for url := range input {
	 result := downloadImage(url)
	 results <- result
	}
   }


func downloadImage(url string) DownloadResult {
	resp, err := http.Get(url)
	if err != nil {
	 return DownloadResult{URL: url, Error: err}
	}
	defer resp.Body.Close()
   
	// Extract the filename from the content disposition header
	fileName, err := getFileNameFromHeader(resp.Header)
	if err != nil {
	 // If content disposition header is not present, generate a filename based on the URL
	 fileName = getFileNameFromURL(url)
	}
   
	// Create a file to save the image
	file, err := os.Create(fileName)
	if err != nil {
	 return DownloadResult{URL: url, Error: err}
	}
	defer file.Close()
   
	// Save the image to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
	 return DownloadResult{URL: url, Error: err}
	}
   
	return DownloadResult{URL: url, Error: nil}
   }
   