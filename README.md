#  Mastering Go Concurrency

This repository is my personal lab for learning and implementing Concurrency in Go. It tracks my progress from basic Goroutines to advanced patterns like Pipelines and Fan-out/Fan-in.

##  Roadmap & Progress

-  **Goroutines**: Basic usage and the `main` goroutine lifecycle.
-  **WaitGroups**: Coordinating multiple goroutines using `sync.WaitGroup`.
-  **Channels**: Communication between goroutines (Buffered vs. Unbuffered).
-  **Select Statement**: Managing multiple channel operations.
-  **Mutexes & RWMutex**: Handling Race Conditions and shared memory.
   **Context Package**: Cancellation, Timeouts, and Deadlines.
-  **Concurrency Patterns**:
    -  Worker Pools
    -  Fan-out / Fan-in
    -  Pipelines
    -  Generators

##  Project Structure

Each folder contains a specific concept or a mini-project:

- `/01-goroutines`: Basics of spawning and the scheduler.
- `/02-waitgroups`: Solving the "main exits too early" problem.
- `/exercises`: Small challenges and solutions.

##  How to Run

To run any of the examples, navigate to the directory and use:

```bash
go run main.go
```

##  Key Learnings
*   **"Do not communicate by sharing memory; instead, share memory by communicating."**
*   The difference between Concurrency (design) and Parallelism (execution).
*   Avoiding goroutine leaks by ensuring every routine has a way to exit.

---
*Follow my journey as I build high-performance Go systems!*
