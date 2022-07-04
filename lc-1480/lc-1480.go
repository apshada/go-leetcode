package main

import "fmt"

func main() {
	nums := [5]int{3, 1, 2, 10, 1}
	result := make([]int, 0)
	for i, number := range nums {
		if i == 0 {
			result = append(result, nums[i])
		}
		if i > 0 {
			result = append(result, result[i-1]+nums[i])
		}

		fmt.Println("OK", i, number, result)
	}
}
