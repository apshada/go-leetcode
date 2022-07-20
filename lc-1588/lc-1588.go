package main

import "fmt"

func main() {
	fmt.Println("Prefix Sum")
	nums := [5]int{1, 4, 2, 5, 3}
	prefixSum := make([]int, 0)

	rightSum := 0
	leftSum := 0
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			prefixSum = append(prefixSum, nums[i])
		} else {
			prefixSum = append(prefixSum, prefixSum[i-1]+nums[i])
			// if i%2 != 0 {
			// 	rightSum = rightSum + (prefixSum[i-1] + nums[i])
			// } else {
			// 	rightSum = rightSum + prefixSum[i-1]
			// }
		}
	}
	// sum := 0
	// for j := 0; j < len(prefixSum); j++ {
	// 	if j%2 != 0 {
	// 		fmt.Println(rightSum-nums[j], rightSum)
	// 		sum = rightSum - nums[j]
	// 	}

	// }
	fmt.Println(nums, prefixSum, rightSum, leftSum)
}
