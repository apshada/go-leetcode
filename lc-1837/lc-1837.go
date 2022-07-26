package main

import "fmt"

func main() {
	n := 10
	k := 10
	count := 0
	for n > 0 {
		remainder := n % k
		count = count + remainder
		n = n / k
	}
	fmt.Println(count)
}
