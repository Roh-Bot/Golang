package main

import "fmt"

func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type person struct {
}

func MyFunc[T comparable](T) {

}

func main() {
	var m = map[int]string{1: "2", 2: "4", 4: "8"}
	fmt.Println("keys:", MapKeys(m))
}
