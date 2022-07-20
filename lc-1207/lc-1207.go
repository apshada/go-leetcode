package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 2, 1, 3}
	res := make(map[int]int, 0)
	res2 := make(map[int]int, 0)
	for i := 0; i < len(arr); i++ {
		res[arr[i]]++
	}
	fmt.Println(res)

	for i, r := range res {
		fmt.Println(i, r)
		// if _, res[r] := res2;
		if _, found := res2[r]; found {
			fmt.Println(false)
		} else {
			res2[r]++
		}
		// if _, res2[r] :=
	}
	fmt.Println(true)

}
