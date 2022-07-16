package main

import "fmt"

func main() {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	low := 0
	high := len(s) - 1
	for low < high {
		s[low], s[high] = s[high], s[low]
		low++
		high--
	}
	fmt.Println(s)
}
