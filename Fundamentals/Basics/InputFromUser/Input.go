package main

import (
	"bufio"
	"fmt"
	"os"
)

var i int

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter rating of our product")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	fmt.Println("Thanks for rating ", input)

}
