package main

import (
	"fmt"
	"net/http"
)

type WebServer bool

func (web WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "text/html")
	fmt.Fprintln(w, "Hello Dhebug")
	//fmt.Fprintf(w, "Request is %v", r)
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
	var webServer1 WebServer
	http.ListenAndServe("localhost:8080", webServer1)
}
