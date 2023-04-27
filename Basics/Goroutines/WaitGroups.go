package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1) // add one goroutine to the WaitGroup

		go func(num int) {
			defer wg.Done() // signal that this goroutine is done when it finishes

			fmt.Println("Goroutine", num, "started")
			// do some work...
			fmt.Println("Goroutine", num, "finished")
		}(i)
	}

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines finished")
}
