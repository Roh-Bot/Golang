package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // signal that this worker is done when it finishes

	fmt.Println("Worker", id, "started")
	// do some work...
	fmt.Println("Worker", id, "finished")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)         // add one worker to the WaitGroup
		go worker(i, &wg) // pass the WaitGroup by pointer
	}

	wg.Wait() // wait for all workers to finish
	fmt.Println("All workers finished")
}
