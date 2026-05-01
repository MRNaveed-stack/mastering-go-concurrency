package main

import (
	"fmt"
	"sync"
)

// This is an example where a race condition will occurr
func main() {
	var wg sync.WaitGroup
	scores := make(map[string]int)

	// We use a large loop to ensure both goroutines are
	// trying to write at the exact same time multiple times.

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			scores["player1"] = 10
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10000; i++ {
			scores["player1"] = 20
		}
	}()
	wg.Wait()
	fmt.Println(scores)

}

// We can detect race condition by command:
// go run -race main.go
