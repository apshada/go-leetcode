package main

import "fmt"

func main() {
	n := "32"
	ans := 0
	for _, v := range n {
		if int(v-'0') > ans {
			fmt.Println(int(v - '0'))
			ans = int(v - '0')
		}
	}
	// fmt.Println(ans)
}
