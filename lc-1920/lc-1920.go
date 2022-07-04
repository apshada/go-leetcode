package main

import "fmt"

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	result := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		result = append(result, nums[nums[i]])
	}
	fmt.Println(result)
}
