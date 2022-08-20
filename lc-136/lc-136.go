package main

import "fmt"

func main() {
	nums := []int{2, 2, 1}
	res := make(map[int]int, 0)
	count := 0
	for _, num := range nums {
		res[num]++
	}

	for key, val := range res {
		if(val == 1){
			count = key
		}
	}
	fmt.Println(count)
}
//Comment