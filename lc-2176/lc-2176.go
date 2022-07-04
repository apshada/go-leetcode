package main

import "fmt"

func main() {
	nums := []int{3, 1, 2, 2, 2, 1, 3}
	k := 2
	countPairs(nums, k)
}

func countPairs(nums []int, k int) int {
	pairs := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && (i*j)%k == 0 {
				fmt.Println(i, j)
				pairs++
			}
		}
	}
	return pairs
}
