package main

import "fmt"

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	res := make(map[int]int, 0)
	c := 0
	for i := 0; i < len(nums); i++ {
		res[nums[i]]++
	}

	for i, _ := range res {
		if res[i] > len(nums)/2 {
			c = i
		}
	}
	fmt.Println(c)
}
