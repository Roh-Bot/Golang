package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		fmt.Println(<-messages)
	}()

	go func() {
		messages <- "Hello Dhebug"
	}()
	msg := <-messages
	fmt.Println(msg)
}
