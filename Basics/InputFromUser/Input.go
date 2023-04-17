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

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating ", input)

}
