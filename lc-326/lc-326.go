package main

import "fmt"

func main() {
	n := 27
	res := isPowerOfThree(n)
	fmt.Println(res)
}

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%3 == 0 {
		return isPowerOfThree(n / 3)
	}
	return false
}
