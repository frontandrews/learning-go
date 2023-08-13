package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d started\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Number of workers
	numWorkers := 3

	// Add workers to the wait group
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	fmt.Println("All workers have finished!")

	// Using an empty struct to signal completion
	done := make(chan struct{})
	go func() {
		fmt.Println("Performing cleanup...")
		time.Sleep(time.Second)
		close(done)
	}()

	// Wait for cleanup to finish
	<-done
	fmt.Println("Cleanup completed!")
}
