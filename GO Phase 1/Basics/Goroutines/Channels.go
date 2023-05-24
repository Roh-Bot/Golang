package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() {
		messages <- "Hello Dhebug"
	}()

	go func() {
		fmt.Println(<-messages)
	}()
	msg := <-messages
	fmt.Println(msg)
}
