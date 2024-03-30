package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Pages  int    `json:"pages"`
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	book := Book{
		"Gaming",
		"Garo",
		25,
	}
	json.NewEncoder(writer).Encode(book)
}

func Hello(write http.ResponseWriter, request *http.Request) {
	write.Header().Set("Content-Type", "text/html")
	write.Write([]byte("<h1 style = 'color : steelblue'>Hello Dhebug</h1>"))
}

func main() {
	//http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
	//	writer.Write([]byte("Hello Dhebug"))
	//})
	http.HandleFunc("/", Hello)
	http.HandleFunc("/books", GetBook)
	log.Fatal(http.ListenAndServe("localhost:5000", nil))

}
