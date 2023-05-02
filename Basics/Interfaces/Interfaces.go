package main

import "fmt"

type Book interface {
	ShowAuthor()
}

func PrintAuthorObject(b Book) {
	fmt.Println(b)
}

type Author struct {
	AuthorName string
}

func (a Author) ShowAuthor() {
	fmt.Println(a.AuthorName)
}

func main() {
	author1 := Author{"Garo"}
	var author2 Author
	author1.ShowAuthor()
	PrintAuthorObject(author1)
	PrintAuthorObject(author2)

}
