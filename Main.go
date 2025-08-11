package main

import (
	"fmt"
	"sync"
)

var urls = []string{
	"https://123.com",
	"https://123.com",
	"https://123.com",
}

func worker(job chan string, res chan string, wg *sync.WaitGroup) {
	for x := range job {
		result := x + "a"
		res <- result
	}
	wg.Done()

}
func main() {
	job := make(chan string, len(urls))
	res := make(chan string, len(urls))

	var wg sync.WaitGroup
	for _, v := range urls {
		wg.Add(1)
		go worker(job, res, &wg)
		job <- v
	}
	close(job)

	wg.Wait()
	close(res)

	for x := range res {
		fmt.Println(x)
	}

}
