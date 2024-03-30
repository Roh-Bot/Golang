package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	i := 0
	for i = 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i, day := range days {
		fmt.Println("Index number:", i, "\nDayPrint:", day)
	}

	for _, day := range days {
		if day == "Friday" {
			goto Dhebug
		}
		fmt.Println(day)
	}

Dhebug:
	fmt.Println("Hello Dhebug")

	//map1 := map[int]string{1:"string",2:"Int"}

}
