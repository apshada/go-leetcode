package main

import "fmt"

func main() {
	n := 3
	res := fizzBuzz(n)
	fmt.Println(res)
}
func fizzBuzz(n int) []string {
	res := make([]string, 0)
	for i := 1; i <= n; i++ {
		if i%5 == 0 && i%3 == 0 {
			res = append(res, "FizzBuzz")
		} else if i%5 == 0 {
			res = append(res, "Fizz")
		} else if i%3 == 0 {
			res = append(res, "Buzz")
		} else {
			s := fmt.Sprintf("%v", i)
			res = append(res, s)
		}
	}
	return res
}
