package engine

import (
	"web-crawler/pkg/fetcher"
	"web-crawler/pkg/frontier"
	"fmt"
	"sync"
)

// Crawl is the main coordinator.
// It takes a seed URL and a max depth to prevent infinite crawling.
func Crawl(seed string, maxDepth int, registry *frontier.Registry) {
	// Channels for the pipeline
	// workQueue: URLs waiting to be fetched
	workQueue := make(chan string, 1000)
	var wg sync.WaitGroup

	// 1. Add the first URL to the queue
	wg.Add(1)
	go func() { workQueue <- seed }()

	// 2. Start the Dispatcher
	// We use a loop to keep the program alive as long as there is work
	fmt.Printf("--- Starting Crawl on %s ---\n", seed)

	// Since we are doing a recursive crawl, we'll launch goroutines
	// dynamically. To keep it safe, we'll limit the "concurrency"
	// using a buffered channel as a semaphore.
	semaphore := make(chan struct{}, 20) // Limit to 20 concurrent HTTP requests

	for i := 0; i < maxDepth; i++ {
		// In a real crawler, depth management is complex.
		// For this version, we will process the queue until it's empty.
		go func() {
			for url := range workQueue {
				// Mark as visited first! (Layer 3)
				if !registry.Visit(url) {
					wg.Done()
					continue
				}

				// Launch a worker
				go func(target string) {
					semaphore <- struct{}{}        // Acquire slot
					defer func() { <-semaphore }() // Release slot
					defer wg.Done()

					fmt.Printf("Fetching: %s\n", target)
					links, err := fetcher.ExtractLinks(target)
					if err != nil {
						return
					}

					// Add new links back to the queue
					for _, link := range links {
						wg.Add(1)
						workQueue <- link
					}
				}(url)
			}
		}()
	}

	wg.Wait()
	close(workQueue)
	fmt.Println("--- Crawl Finished ---")
}
