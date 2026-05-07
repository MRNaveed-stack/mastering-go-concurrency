package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

type Task struct {
	ID    int
	Data  string
	State string
}

func fetchTaskIDs(ctx context.Context, count int) <-chan Task {
	out := make(chan Task)
	go func() {
		defer close(out)
		for i := 1; i <= count; i++ {
			select {
			case out <- Task{ID: i}:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func downloadImages(ctx context.Context, in <-chan Task, limit time.Duration) <-chan Task {
	out := make(chan Task)
	ticker := time.NewTicker(limit)

	go func() {
		defer close(out)
		defer ticker.Stop()
		for task := range in {
			select {
			case <-ticker.C:
				task.Data = fmt.Sprintf("Image-Data-%d", task.ID)
				task.State = "Downloaded"

				select {
				case out <- task:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func processImages(ctx context.Context, in <-chan Task) <-chan Task {
	out := make(chan Task)
	go func() {
		defer close(out)
		for task := range in {
			time.Sleep(time.Duration(rand.IntN(100)) * time.Millisecond)
			task.State = "Processed/Filtered "

			select {
			case out <- task:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	fmt.Println("System initializing pipeline")
	source := fetchTaskIDs(ctx, 20)
	download := downloadImages(ctx, source, 200*time.Millisecond)

	processed := processImages(ctx, download)
	count := 0
	for result := range processed {
		count++
		fmt.Printf("Completed Task %d: , %s", result.ID, result.State)
	}
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Pipeline timed out")
	} else {
		fmt.Println("All task completed successfully")
	}
}
