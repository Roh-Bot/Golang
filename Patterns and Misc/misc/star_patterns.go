package misc

import "fmt"

func RectangleWithDiagonals(k int) {
	// Nested 2 for loops for Matrix printing
	// Outer loop for rows
	for a := 1; a <= k; a++ {
		// Inner loop for columns
		for b := 1; b <= k; b++ {
			// Condition check using OR operator where
			// 1. Printing star as per first 4 checks
			// on the circumference of rectangle
			// 2. Fifth check is for diagonals
			if a == 1 || a == k || b == 1 || b == k || a == b || b == (k-a+1) {
				fmt.Print("*")

			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
