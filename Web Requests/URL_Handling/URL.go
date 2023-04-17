package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbjghb"

func main() {
	//response, err := http.Get(myurl)
	//
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(response)
	//response.Body.Close()

	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	var qparams url.Values = result.Query()
	fmt.Printf("The type of qparams is: %T\n", qparams)
	fmt.Println(qparams["coursename"])
}
