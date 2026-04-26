package main

import (
	"fmt"
	"time"
)

func sayHello(name string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello %s (Iteration %d)\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func sayhi(name string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("Hi %s (Iteration %d)\n", name, i)
		time.Sleep(100 * time.Millisecond)
	}
}
func main() {
	go sayHello("Goroutine")
	fmt.Println("Main function calling...")
	sayhi("Main Thread")

}
