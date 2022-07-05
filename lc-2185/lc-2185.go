package main

import "fmt"

func main() {
	words := []string{"pay", "attention", "practice", "attend"}
	pref := "at"

	result := 0

	for _, word := range words {
		if len(pref) <= len(word) && word[:len(pref)] == pref {
			result++
		}
	}

	fmt.Println(result)
}
