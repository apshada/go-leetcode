package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	num := 2932
	snum := strconv.Itoa(num)
	arr := make([]int, 0)
	for _, n := range snum {
		if int(n-'0') != 0 {
			arr = append(arr, int(n-'0'))
		}
		// fmt.Println(i, string(n))
	}
	sort.Ints(arr)
	fmt.Println(arr)
	fmt.Println(arr[0] + arr[1])
	// arr = append(arr, num/100)
	// arr = append(arr, num/1000)
	// arr = append(arr, num % 10)
	// fmt.Println(arr)
}
