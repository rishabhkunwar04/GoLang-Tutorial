# Golang Convepts


### Worker pool pattern

```Go
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Simulated list of URLs to download
var urls = []string{
	"https://example.com/a",
	"https://example.com/b",
	"https://example.com/c",
	"https://example.com/d",
	"https://example.com/e",
	"https://example.com/f",
}

// Simulated download function
func download(url string) string {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return fmt.Sprintf("Downloaded: %s", url)
}

// Worker function
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		fmt.Printf("Worker %d started %s\n", id, url)
		result := download(url)
		fmt.Printf("Worker %d finished %s\n", id, url)
		results <- result
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	numWorkers := 3
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// Send jobs
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	// Wait for all workers
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Println("Result ->", result)
	}
}

```