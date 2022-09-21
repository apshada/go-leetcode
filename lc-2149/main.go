package main

import "fmt"

func main() {
	// arr := []int{3, 1, -2, -5, 2, -4}
	nums := []int{-1, 1}
	posArr := make([]int, 0)
	negArr := make([]int, 0)
	res := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			posArr = append(posArr, nums[i])
		} else {
			negArr = append(negArr, nums[i])
		}
	}

	for j := 0; j < len(nums)/2; j++ {
		res = append(res, posArr[j])
		res = append(res, negArr[j])
	}

	fmt.Println(res)

}
