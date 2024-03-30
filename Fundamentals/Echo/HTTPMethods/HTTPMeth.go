package main

import "net/http"

func myLogin(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("content-type", "text/html")
		w.Write([]byte(`
	<html>
			<head>
				Hi
			</head>
			<body>
				<h1 style = "color:blue">
					Hello Dhebug GET
				</h1>
			</body>
	</html>
`))
	}
	if r.Method == http.MethodPost {
		w.Header().Set("content-type", "text/html")
		w.Write([]byte(`
	<html>
			<head>
				Hi
			</head>
			<body>
				<h1 style = "color:blue">
					Hello Dhebug POST	
				</h1>
			</body>
	</html>
`))
	}

}

func main() {
	http.HandleFunc("/login", myLogin)
	http.ListenAndServe("localhost:8080", nil)
}
