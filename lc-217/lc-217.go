package main

import "fmt"

func main() {
	nums := [4]int{1, 2, 3, 1}
	res := make(map[int]int, 0)

	for _, nums := range nums {
		res[nums] = res[nums] + 1
		fmt.Println(res[nums])
		if res[nums] > 1 {
			// return true
		}
	}
	// fmt.Println("")
	// return false
}
