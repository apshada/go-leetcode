package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	res := make([]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		// fmt.Println(i)
		x := nums[i]
		y := nums[i+1]
		for x > 0 {
			res = append(res, y)
			x--
		}
		fmt.Println(res, i)
		i = i + 1
	}

}
