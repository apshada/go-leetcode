package main

import "fmt"

func main() {
	gain := []int{-5, 1, 5, 0, -7}
	largestAltitude(gain)
}

func largestAltitude(gain []int) int {
	prefixSum := make([]int, 0)

	for i, _ := range gain {
		if i == 0 {
			prefixSum = append(prefixSum, 0)
			prefixSum = append(prefixSum, gain[i])
		} else {
			prefixSum = append(prefixSum, prefixSum[i]+gain[i])
		}
	}
	max := -100
	for i, _ := range prefixSum {
		if prefixSum[i] > max {
			max = prefixSum[i]
		}
	}
	fmt.Println(max)
	return max
}
