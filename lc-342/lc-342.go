package main

import "fmt"

func main() {
	n := 16
	res := isPowerOfFour(n)
	fmt.Println(res)
}

func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%4 == 0 {
		return isPowerOfFour(n / 4)
	}
	return false
}
