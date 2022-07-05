package main

import (
	"fmt"
)

func main() {
	boxes := "001011"
	minsops, left, right := 0, 0, 0
	for i, box := range boxes {
		// j := 0
		if string(box) == "1" {
			minsops += i
		}
	}

	ops := make([]int, len(boxes))

	for i, box := range boxes {
		ops[i] = minsops
		right -= int(box - '0')
		fmt.Println(ops[i], box, minsops, right, left)
		// left += int(box - '0')
		// minsops -= right
		// minsops += left
	}
	fmt.Println(ops)
}
