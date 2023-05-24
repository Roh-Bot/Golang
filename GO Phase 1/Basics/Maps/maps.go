package main

import (
	"fmt"
)

func main() {
	//map1 := map[int]string{1: "G", 2: "Dhebug", 3: "Garo"}

	//Map of integer to a slice of string
	map1 := map[int][]string{
		1: {"DHebug", "bhai"},
		2: {"", "", ""},
	}
	fmt.Println(map1)

	map2 := make(map[int]string)
	map2[1] = "G"
	map2[2] = "Dhebug"
	map2[3] = "Garo"

	delete(map2, 3)
	fmt.Println(map2)

	for key, value := range map2 {
		fmt.Printf("Key = %v \nValue = %v\n", key, value)
	}

	for _, value := range map2 {
		fmt.Printf("Key v \nValue = %v\n", value)
	}

}
