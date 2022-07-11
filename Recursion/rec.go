package main

import "fmt"

func main() {
	// printName(0, 5)
	// linearPrint(0, 10)
	linearPrintUsingBacktrack(10, 10)
	// NonlinearPrint(10)
}

func printName(i, n int) {
	if i != n {
		fmt.Println("Aditya")
		printName(i+1, n)
	}
	return
}

func linearPrint(i, n int) {
	if i != n {
		fmt.Println(i)
		linearPrint(i+1, n)
	}
	return
}

func NonlinearPrint(n int) {
	if n > 0 {
		fmt.Println(n - 1)
		NonlinearPrint(n - 1)
	}
	return
}

func linearPrintUsingBacktrack(i, n int) {
	if i < 1 {
		return
	}
	linearPrintUsingBacktrack(i-1, n)
	fmt.Println(i)
}
