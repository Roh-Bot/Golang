package main

import (
	"fmt"
	"net/http"
)

func myHandler2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `
	<html>
		<head>
			Hi
		</head>
		<body>
			<h1 style = "color:blue">
				Hello Dhebug
			</h1>
        </body>
	</html>
`)
}

func main() {
	http.Handle("/hello", http.HandlerFunc(myHandler2))
	http.ListenAndServe("localhost:8080", nil)
}
