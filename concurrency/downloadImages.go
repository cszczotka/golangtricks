package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

func extractFileName(url string) string {
	// Extract the file name from the URL (assuming the last segment of the path is the file name)
	parts := strings.Split(url, "/")
	return parts[len(parts)-1] + ".jpg"
}

func downloadImage(url string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// Fetch image from the URL
	resp, err := http.Get(url)
	if err != nil {
		results <- fmt.Sprintf("Error downloading %s: %s", url, err)
		return
	}
	defer resp.Body.Close()

	// Extract the file name from the URL
	fileName := extractFileName(url)

	// Create a file to save the image
	file, err := os.Create(fileName)
	if err != nil {
		results <- fmt.Sprintf("Error creating file for %s: %s", url, err)
		return
	}
	defer file.Close()

	// Save the image to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		results <- fmt.Sprintf("Error saving %s: %s", url, err)
		return
	}

	results <- fmt.Sprintf("Downloaded %s", url)
}

func main() {
	var wg sync.WaitGroup
	results := make(chan string, 3)

	imageURLs := []string{"https://picsum.photos/id/237/200/300", "https://picsum.photos/seed/picsum/200/250", "https://picsum.photos/200/300?grayscale"}

	// Fan-out: Start goroutines to download each image concurrently
	for _, url := range imageURLs {
		wg.Add(1)
		go downloadImage(url, &wg, results)
	}

	// Fan-in: Collect results from goroutines
	go func() {
		wg.Wait()
		close(results)
	}()

	// Wait for all goroutines to finish and print results
	for result := range results {
		fmt.Println(result)
	}
}
