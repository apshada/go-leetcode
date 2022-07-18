package main

import "fmt"

func main() {
	encoded := []int{6, 2, 7, 3}
	first := 4
	flag := 0
	result := make([]int, 0)
	result = append(result, first)
	for i, res := range encoded {
		if i == 0 {
			flag = res ^ first
			result = append(result, flag)
		} else {
			flag = flag ^ res
			result = append(result, flag)
		}
	}
	fmt.Println(result)
}
