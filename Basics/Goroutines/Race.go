package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Race Condition")

	wg := &sync.WaitGroup{}
	mut := &sync.Mutex{}

	var score = []int{0}
	wg.Add(4)
	go func(wgParam *sync.WaitGroup, mutParam *sync.Mutex) {
		defer wgParam.Done()
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mutParam.Unlock()
	}(wg, mut)
	go func(wgParam *sync.WaitGroup, mutParam *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)
	go func(wgParam *sync.WaitGroup, mutParam *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		defer wg.Done()
		fmt.Println("Score R")
		mut.Lock()
		fmt.Println(score)
		mut.Unlock()
	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}

//package main
//
//import (
//"fmt"
//"sync"
//)
//
//func main() {
//	fmt.Println("Race Condition")
//
//	wg := &sync.WaitGroup{}
//	mut := &sync.Mutex{}
//
//	var score = []int{0}
//	wg.Add(2)
//	go func() {
//		defer wg.Done()
//		fmt.Println("One R")
//		mut.Lock()
//		score = append(score, 1)
//		mut.Unlock()
//	}()
//	go func() {
//		defer wg.Done()
//		fmt.Println("Two R")
//		mut.Lock()
//		score = append(score, 2)
//		mut.Unlock()
//	}()
//	wg.Wait()
//	fmt.Println(score)
//}
