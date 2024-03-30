package main

import (
	"fmt"
	"strings"
)

func main() {
	s := []string{"1", "2", "3", "4", "5"}
	fmt.Println(strings.Join(s, ", "))
}
