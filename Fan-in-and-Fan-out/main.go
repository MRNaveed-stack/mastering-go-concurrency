package main

import (
	"context"
	"fmt"
	"sync"
)

// This is our producer
// Our current approach can cause Goroutine leak if there is no reciever channel at the other end
// We have a solution for this implemented in the same function but modifying this function a little bit
// func generator(nums ...int) <-chan int {
// 	out := make(chan int)
// 	go func() {
// 		for _, n := range nums {
// 			out <- n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

// This function will ensure no goroutine will be leaked
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
				// successfully sent

			case <-done:
				return
			}
		}
	}()
	return out
}

// This will be fanned out
// let's apply context.Context to prevent our program from crashing with big and unexpected data
// we will apply context.Context in the copy of this square function
// func square(in <-chan int) <-chan int {

// 	out := make(chan int)
// 	go func() {
// 		for n := range in {
// 			out <- n * n
// 		}
// 		close(out)
// 	}()
// 	return out
// }

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
				//Normal work
			case <-ctx.Done():
				return
			}
		}
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
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	done := make(chan struct{})
	defer close(done)

	in := generator(done, 1, 2, 3, 4, 5, 6, 7, 8)

	worker1 := square(ctx, in)
	worker2 := square(ctx, in)
	worker3 := square(ctx, in)

	for n := range merge(worker1, worker2, worker3) {
		fmt.Printf("%d", n)
	}
}
