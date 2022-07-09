// Golang program to illustrate the
// string matching using regexp
// in-built functions
package main

import (
	"fmt"
)

func main() {
	x := 0
	ad := []string{"--X", "X++", "X++"}
	for i := range ad {
		if ad[i] == "--X" || ad[i] == "X--" {
			x = x - 1
		} else if ad[i] == "++X" || ad[i] == "X++" {
			// (ad[i] == "--X")
			x = x + 1
		}
	}
	fmt.Println(x)
}
