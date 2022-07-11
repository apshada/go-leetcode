package main

import "fmt"

func main() {
	n := 16
	res := isPowerOfTwo(n)
	fmt.Println(res)
}

func isPowerOfTwo(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	if n%4 == 0 {
		return isPowerOfTwo(n / 4)
	}
	return false
}

// 0 1 2 3 4  5  6
// 1 4 16 8 16 32 64
