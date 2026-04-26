package main

// Unbuffered channel example
// func main() {
// 	ch := make(chan string)
// 	go func() {
// 		fmt.Println("Sending message...")
// 		time.Sleep(2 * time.Second)
// 		ch <- "Hello this is the message i wanted to send"
// 		fmt.Println("Message sent.")

// 	}()
// 	fmt.Println("Waiting for message...")
// 	msg := <-ch
// 	fmt.Printf("Received message: %s\n", msg)

// }

// Buffered channel example
// func main() {

// 	ch := make(chan int, 2)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println("I wanna send a message")
// 	// ch <- 3
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)

// }

// Unidirectional channels
// func produce(p chan<- int) {
// 	p <- 42
// 	close(p)
// }

// func consume(a <-chan int) {
// 	fmt.Println(<-a)
// }

// func main() {
// 	ch := make(chan int)
// 	go produce(ch)
// 	consume(ch)
// }

// Select keyword usage
// when we have two go routines then whatever goroutine takes less time will get executed
// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch1 <- "Data from channel 1"

// 	}()

// 	go func() {
// 		time.Sleep(1 * time.Second)
// 		ch2 <- "Data from channel 2"
// 	}()

// 	select {
// 	case rec1 := <-ch1:
// 		fmt.Println("Recieved from channel 1", rec1)
// 	case rec2 := <-ch2:
// 		fmt.Println("Recieved from channel 2", rec2)
// 	}

// }

// Select usage with timeout . Let's simulate it
// func main() {
// 	ch := make(chan string)

// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		ch <- "Database result"
// 	}()

// 	select {
// 	case rec1 := <-ch:
// 		fmt.Println("Success: ", rec1)
// 	case <-time.After(2 * time.Second):
// 		fmt.Println("Operation timeout")

// 	}
// }

// Non blocking operations with default
// func main() {
// 	message := make(chan string)
// 	select {
// 	case msg := <-message:
// 		fmt.Println("Message recieved is: ", msg)
// 	default:
// 		fmt.Println("Going to main because there was nothing in the channel that was sent")
// 	}
// }
