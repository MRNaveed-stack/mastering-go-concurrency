package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// This is an example of how we will use sync.once

// type Database struct {
// 	connection string
// }

// var (
// 	dbInstance *Database
// 	once       sync.Once
// )

// func GetDatabase() *Database {
// 	once.Do(func() {
// 		fmt.Println("Connecting to the Database (Extensive operation)")
// 		dbInstance = &Database{connection: "Postgres connected"}
// 	})
// 	return dbInstance
// }

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 1; i <= 10; i++ {
// 		wg.Add(1)
// 		go func(id int) {
// 			defer wg.Done()
// 			db := GetDatabase()
// 			fmt.Printf("Worker %d using: %s\n", id, db.connection)
// 		}(i)
// 	}
// 	wg.Wait()
// }

// This is an example of how we will use sync.Atomic for atomic operations
func main() {
	var (
		counter atomic.Int64
		wg      sync.WaitGroup
	)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Add(1)
		}()
	}
	wg.Wait()
	fmt.Println("Final Counter Value: ", counter.Load())
}
