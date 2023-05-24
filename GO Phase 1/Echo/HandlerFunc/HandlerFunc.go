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
				Hello Dhebug Welcome
			</h1>
        </body>
	</html>
`)
}

func main() {
	http.HandleFunc("/welcome", myHandler)
	http.ListenAndServe("localhost:8080", nil)
}
