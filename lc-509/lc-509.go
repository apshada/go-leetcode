package main

import "fmt"

func main() {
	n := 4
	res := fib(n)
	fmt.Println(res)
}

func fib(n int) int {
	if n == 0 {
		return n
	}
	if n == 1 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

// 0 1 1 2
