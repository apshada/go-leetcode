package main

import "fmt"

func main() {
	// Prefix Sum
	nums := [7]int{2, 3, -1, 8, 4, 4, 5}
	// nums := [3]int{1, -1, 4}
	prefixSum := make([]int, 0)
	prefixSum2 := make([]int, 0)
	prefixSum = append(prefixSum, nums[0])
	prefixSum2 = append(prefixSum2, nums[len(nums)-1])
	for i, _ := range nums {
		if i > 0 {
			prefixSum = append(prefixSum, prefixSum[i-1]+nums[i])
		}
	}

	for j := len(nums) - 1; j < 0; j-- {
		prefixSum2 = append(prefixSum2, prefixSum2[j-1]+nums[j])
	}
	fmt.Println(prefixSum)
	fmt.Println(prefixSum2)
}
