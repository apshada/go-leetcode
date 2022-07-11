package main

import (
	"fmt"
	"math"
)

func main() {
	a := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	max_ending_here := 0
	max_so_far := math.MinInt64

	for i := 0; i < len(a); i++ {
		max_ending_here = max_ending_here + a[i]
		if max_so_far < max_ending_here {
			max_so_far = max_ending_here
		}
		if max_so_far < 0 {
			max_ending_here = 0
		}
	}
	fmt.Println(max_so_far)
}
