package main

import "fmt"

// "fmt"
// "sync"
// "time"

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting ...\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Worker %d finished!\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 3; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}
// 	fmt.Println("Main is waiting for workers...")
// 	wg.Wait()
// 	fmt.Println("All workers done. Proceeding with main...")
// }

// func main() {
// 	var wg sync.WaitGroup
// 	names := []string{"Alice", "Bob", "Charlie"}
// 	for _, name := range names {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Printf("Hello, %s!\n", name)
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("All greetings done.")
// }

// Another example
// func processTask(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("Worker %d starting heavy task...\n", id)
// 	time.Sleep(time.Second)
// 	fmt.Printf("Task %d is completed successfully", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	tasks := 3
// 	for i := 1; i <= tasks; i++ {
// 		wg.Add(1)
// 		go processTask(i, &wg)
// 	}
// 	fmt.Println("Main: Waiting for all goroutines to finish")
// 	wg.Wait()
// 	fmt.Println("All workers finished moving to the next step")
// }

// func fetchSource(name string, wg *sync.WaitGroup, ch chan<- string) {
// 	defer wg.Done()
// 	time.Sleep(time.Second)
// 	ch <- fmt.Sprintf("%s data retrieved", name)

// }

// func main() {
// 	var wg sync.WaitGroup
// 	results := make(chan string)
// 	sources := []string{"User", "Orders", "Notifications"}
// 	start := time.Now()
// 	for _, s := range sources {
// 		wg.Add(1)
// 		go fetchSource(s, &wg, results)
// 	}
// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()
// 	for res := range results {
// 		fmt.Println(res)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// 	fmt.Printf("Total time taken: %v\n", time.Since(start))
// }

// Worker Pool
// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for j := range jobs {
// 		fmt.Printf("Worker %d processing job %d\n", id, j)
// 		time.Sleep(500 * time.Millisecond)
// 		results <- j * j
// 	}
// }
// func main() {
// 	const numJobs = 10
// 	jobs := make(chan int, numJobs)
// 	results := make(chan int, numJobs)

// 	for w := 1; w <= 3; w++ {
// 		go worker(w, jobs, results)
// 	}
// 	for j := 1; j <= numJobs; j++ {
// 		jobs <- j
// 	}
// 	close(jobs)

// 	for r := 1; r <= numJobs; r++ {
// 		fmt.Println("Result: ", <-results)
// 	}
// }

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

func main() {
	gen := generator(2, 3, 4, 5)
	sq := square(gen)

	for n := range sq {
		fmt.Println("Final Result: ", n)
	}
}
