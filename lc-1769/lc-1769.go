package main

import (
	"fmt"
)

func main() {
	boxes := "001011"

	for i, box := range boxes {
		j := 0
		if string(box) == "1" {
			j = j + (i - j)
			fmt.Println(j)
		}
	}
}
