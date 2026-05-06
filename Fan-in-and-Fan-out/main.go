package main

import (
	"fmt"
	"sync"
)

// This is our producer
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// This will be fanned out
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in := generator(1, 2, 3, 4, 5, 6, 7, 8)

	worker1 := square(in)
	worker2 := square(in)
	worker3 := square(in)

	for n := range merge(worker1, worker2, worker3) {
		fmt.Printf("%d", n)
	}
}
