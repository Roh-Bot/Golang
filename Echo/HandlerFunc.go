package main

import (
	"fmt"
	"net/http"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
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
	http.ListenAndServe("localhost:8080", http.HandlerFunc(myHandler))
}
