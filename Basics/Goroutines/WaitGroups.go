package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg *sync.WaitGroup
	wg = &sync.WaitGroup{}
	//wg := &sync.WaitGroup{}
	wg.Add(2) // add one goroutine to the WaitGroup

	go func() int {
		defer wg.Done()
		fmt.Println("Function 1:", 12*1000)
		return 0
	}()

	go func() int {
		defer wg.Done()
		fmt.Println("Function 2:", 13*29999999)
		return 0
	}()

	wg.Wait() // wait for all goroutines to finish
	fmt.Println("All goroutines finished")
}
