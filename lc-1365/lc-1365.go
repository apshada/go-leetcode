package main

import "fmt"

func main() {
	nums := [5]int{8, 1, 2, 2, 3}
	// fmt.Println(nums)
	res := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		res[nums[i]]++
		// for j := i + 1; j < len(nums); j++ {
		// 	if nums[j] < nums[i] && j != i {
		// 		res[nums[j]]++
		// 	}
		// }
	}
	fmt.Println(res)
}
