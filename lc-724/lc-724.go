package main

import (
	"fmt"
)

func main() {
	// nums := [6]int{1, 7, 3, 6, 5, 6}
	nums := [3]int{2, 1, -1}
	prefixSum := make([]int, 0)
	rightSum := 0
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixSum = append(prefixSum, nums[0])
		} else {
			prefixSum = append(prefixSum, prefixSum[i-1]+nums[i])
			rightSum = prefixSum[i-1] + nums[i]
		}
	}
	for i := 0; i < len(nums); i++ {
		leftSum = leftSum + nums[i]
		// fmt.Println(i, leftSum, rightSum)
		if leftSum == rightSum {
			fmt.Println(i)
		}
		rightSum = rightSum - nums[i]
	}

	fmt.Println(-1)
	// fmt.Println(nums)
	// fmt.Println(prefixSum, rightSum, leftSum)
}
