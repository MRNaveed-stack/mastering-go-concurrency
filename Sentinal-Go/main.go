package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

type Result struct {
	URL          string
	StatusCode   int
	ResponseTime time.Duration
	Err          error
}

type Metrics struct {
	mu          sync.Mutex
	statusCodes map[int]int
	totalChecks atomic.Uint64
}

func checkURL(id int, urls <-chan string, result chan<- Result, metrics *Metrics) {
	for url := range urls {
		start := time.Now()
		resp, err := http.Get(url)
		duration := time.Since(start)
		res := Result{URL: url, ResponseTime: duration, Err: err}
		if err == nil {
			res.StatusCode = resp.StatusCode
			resp.Body.Close()
		}
		metrics.totalChecks.Add(1)
		metrics.mu.Lock()
		metrics.statusCodes[res.StatusCode]++
		metrics.mu.Unlock()
		result <- res
	}

}

func main() {
	urls := []string{
		"https://google.com", "https://github.com",
		"https://golang.org", "https://amazon.com",
		"https://httpbin.org/status/404", "https://httpbin.org/delay/2",
	}
	urlChan := make(chan string, len(urls))
	resChan := make(chan Result, len(urls))

	metrics := &Metrics{statusCodes: make(map[int]int)}
	var wg sync.WaitGroup
	numworkers := 3
	for i := 1; i <= numworkers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			checkURL(id, urlChan, resChan, metrics)
		}(i)
	}
	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)

	go func() {
		wg.Wait()
		close(resChan)
	}()

	fmt.Println("Individual Results")
	for res := range resChan {
		if res.Err != nil {
			fmt.Printf("%s | Error %v\n", res.URL, res.Err)
		} else {
			fmt.Printf("%s | Code: %d | Time: %v\n", res.URL, res.StatusCode, res.ResponseTime)

		}
	}
	fmt.Println("Global Metrics")
	fmt.Printf("Total checks performed: %d\n", metrics.totalChecks.Load())
	fmt.Println("Status Code breakdown")
	for code, count := range metrics.statusCodes {
		fmt.Printf(" [%d]: %d times\n", code, count)
	}
}
