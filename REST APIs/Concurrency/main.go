package main

//
//import (
//	"fmt"
//	"sync"
//)
//
//func main() {
//
//	num1 := make(chan int)
//	num2 := make(chan int)
//
//	go Number1(num1)
//	go Number2(num2)
//
//	val1 := <-num1
//	val2 := <-num2
//
//	fmt.Println(val2, val1)
//}
//
//func Number1(num chan int) {
//	var mutex sync.Mutex
//	mutex.Lock()
//	defer mutex.Unlock()
//	num <- 1
//	return
//}
//
//func Number2(num chan int) {
//	var mutex sync.Mutex
//	mutex.Lock()
//	defer mutex.Unlock()
//	num <- 2
//	return
//}
