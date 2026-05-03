package main

import (
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"word-counter/pkg/counter"
)

func main() {
	start := time.Now()
	fileQueue := make(chan string, 100)
	stats := counter.NewStats()
	var wg sync.WaitGroup

	for i := 1; i <= 4; i++ {
		wg.Add(1)
		go counter.FileWorker(i, fileQueue, stats, &wg)

	}

	files, _ := filepath.Glob("data/*.txt")
	for _, f := range files {
		fileQueue <- f
	}
	close(fileQueue)
	wg.Wait()
	fmt.Println("Word count results")
	for word, count := range stats.Counts() {
		if count > 0 {
			fmt.Printf("%s: %d\n", word, count)
		}
	}
	fmt.Printf("\nDone in: %v\n", time.Since(start))
}
