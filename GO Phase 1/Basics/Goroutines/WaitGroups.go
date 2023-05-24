package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg *sync.WaitGroup
	wg = &sync.WaitGroup{}
	//wg := &sync.WaitGroup{}
	wg.Add(3) // add one goroutine to the WaitGroup

	go func() int {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		return 0
	}()

	go func() int {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
		return 0
	}()

	go func() int {
		defer wg.Done()
		fmt.Println("Function 3:", 0)
		return 0
	}()

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines finished")
}
