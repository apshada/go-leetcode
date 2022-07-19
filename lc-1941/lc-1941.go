package main

import "fmt"

func main() {
	s := "aabbcc"
	res := make(map[string]int, 0)

	for i := 0; i < len(s); i++ {
		res[string(s[i])]++
	}
	flag := 0
	j := 0
	for _, r := range res {
		if j == 0 {
			flag = r
			j++
		}
		if flag != r {
			fmt.Println(false)
		}
	}
	fmt.Println(true)
}
