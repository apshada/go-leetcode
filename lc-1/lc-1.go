package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums); j++ {
			// fmt.Println(nums[i], nums[j])
			if nums[i]+nums[j] == target {
				fmt.Println(i, j)
				res = append(res, i)
				res = append(res, j)
				// return res/
			}
		}
	}
}
