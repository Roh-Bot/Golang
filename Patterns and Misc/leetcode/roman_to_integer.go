package leetcode

import "log"

func RomanToIntger(s string) int {
	r := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	for i := 0; i < len(s)-1; i++ {
		//a := s[i]
		log.Println(r)
	}
	return 0
}
