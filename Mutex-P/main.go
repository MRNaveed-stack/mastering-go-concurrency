package main

import (
	"fmt"
	"sync"
)

// Simple Example of only Mutex

// type SafeCounter struct {
// 	mu    sync.Mutex
// 	value int
// }

// func (c *SafeCounter) Inc(wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	c.mu.Lock()
// 	// Do not try to perform heavy I/O while holding a lock because it can freeze all Goroutines.
// 	c.value++
// 	c.mu.Unlock()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	count := SafeCounter{}

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		// I am just using go keyword to practice Go routines otherwise for this specific example we do not need keyword go behind it.
// 		// So do not use Go keyword until you actually need it. Keep it simple and easy.
// 		go count.Inc(&wg)

// 	}
// 	wg.Wait()
// 	fmt.Println("Final Value:", count.value)

// }

// A little tricky example of Mutex
type Result struct {
	URL    string
	Status int
}

type ScraperReport struct {
	mu      sync.Mutex
	results map[string]int
}

func Worker(id int, jobs <-chan string, report *ScraperReport, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		statusCode := 200

		report.mu.Lock()
		report.results[url] = statusCode
		fmt.Printf("Worker %d saved status for %s\n", id, url)
		report.mu.Unlock()
	}
}

func main() {
	report := &ScraperReport{
		results: make(map[string]int),
	}
	jobs := make(chan string, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go Worker(w, jobs, report, &wg)
	}
	urls := []string{"google.com", "github.com", "golang.org", "medium.com"}
	for _, url := range urls {
		jobs <- url
	}
	close(jobs)

	wg.Wait()
	fmt.Println("Final report")
	for url, status := range report.results {
		fmt.Printf("%s: %d\n", url, status)
	}
}
