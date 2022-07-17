package main

import "fmt"

func main() {
	s := "RLRRLLRLRL"
	a := 0
	b := 0
	c := 0
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "R" {
			a++
		} else if string(s[i]) == "L" {
			b++
		}
		if a == b {
			a = 0
			b = 0
			c++
		}
	}
	// return c
	fmt.Println(c)
}
