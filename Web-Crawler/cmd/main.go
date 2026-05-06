package main

import (
	"web-crawler/pkg/engine"
	"web-crawler/pkg/frontier"
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// Initialize our Layer 3 Registry
	registry := frontier.NewRegistry()

	// Config
	seedURL := "https://golang.org" // Change this to any site you want to test
	maxDepth := 2                   // Be careful with high numbers!

	// Start the Engine
	engine.Crawl(seedURL, maxDepth, registry)

	// Final Summary
	fmt.Printf("\n--- Final Report ---\n")
	fmt.Printf("Unique URLs Found: %d\n", registry.Count())
	fmt.Printf("Time Elapsed: %v\n", time.Since(start))
}
