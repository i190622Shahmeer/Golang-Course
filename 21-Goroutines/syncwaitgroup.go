package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	// fmt.Println("Explore goroutines in Go!")
	fmt.Printf("Worker %d started\n", i)
	//some task is happening
	fmt.Printf("Worker %d finished\n", i)
	defer wg.Done() // Decrement the WaitGroup counter when the goroutine completes
}

func testwaitgroup() {

	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // Increment the WaitGroup counter
		go worker(i, &wg)

	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("All workers done!")
}
