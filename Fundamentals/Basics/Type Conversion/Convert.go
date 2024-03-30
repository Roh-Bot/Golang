package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter rating of our product")

	input, _ := reader.ReadString('\n')

	fmt.Printf("Thanks for rating %T", input)

	numrating, err := strconv.ParseInt(strings.TrimSpace(input), 10 , 64)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Rating is", numrating + 1)
	}

}