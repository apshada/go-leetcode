package main

import "fmt"

func main() {
	nums := [5]int{8, 1, 2, 2, 3}
	resData := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		c := 0
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] && j != i {
				c++

			}
		}
		resData = append(resData, c)
	}
	fmt.Println(resData)
}
