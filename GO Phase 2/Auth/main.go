package main

import (
	"authentication/controllers"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/signin", controllers.Signin)
	http.HandleFunc("/callback", controllers.Callback)
	http.ListenAndServe(":8000", nil)

}
