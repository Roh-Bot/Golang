package main

import (
	"fmt"
)

type Book struct {
	Author string
	Title  string
}

func modifyBook(book *Book) {
	book.Title = "Harry Potter and the Chamber of Secrets"
}

func DisplayError(returnNil bool) (*Book, error) {
	if returnNil {
		return nil, fmt.Errorf("Could not retrieve book")
	}
	book := &Book{
		Author: "J.R.R. Tolkien",
		Title:  "The Lord of the Rings",
	}
	return book, nil
}
func main() {
	book1 := Book{
		Author: "J.K. Rowling",
		Title:  "Harry Potter and the Philosopher's Stone",
	}
	fmt.Printf("Book 1 before modification: %+v\n", book1)

	// Scenario 1: Allowing modification of the original data
	modifyBook(&book1)
	fmt.Printf("Book 1 after modification: %+v\n", book1)

	// Scenario 2: Allowing for nil values
	book3, err := DisplayError(false)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Book 3: %+v\n", book3)
	}
	book4, err := DisplayError(true)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Book 4: %+v\n", book4)
	}
}
