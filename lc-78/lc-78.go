package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	// res := [][]int{}``
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			for k := i; k <= j; k++ {

				fmt.Println(nums[i], nums[j])
			}
		}
	}
}

// i = 1 j = 1, 2, 3
// i = 2 j = 2,3
// i = 3 j = 3
