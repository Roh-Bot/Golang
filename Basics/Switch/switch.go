package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please enter choice and press Enter")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, _ := strconv.Atoi(input)

	fmt.Printf("%T", choice)

	switch choice {
	case 1:
		fmt.Println("Choice is 1")
		fallthrough
	case 2:
		fmt.Println("Choice is 2")
		fallthrough
	case 3:
		fmt.Println("Choice is 3")
	case 4:
		fmt.Println("Choice is 4")
	default:
		fmt.Println("Invalid choice")
	}
}
