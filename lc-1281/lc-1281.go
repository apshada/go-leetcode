package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 234
	stringNum := strconv.Itoa(n)
	sum := 0
	product := 1
	for _, num := range stringNum {
		sum = sum + int(num-'0')
		product = product * int(num-'0')
	}

	fmt.Println(product - sum)
}
