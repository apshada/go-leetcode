package main

import "fmt"

func main() {
	s := "ababcbacadefegdehijhklij"
	res := make(map[string]int)

	for _, str := range s {
		res[string(str)]++
	}

	fmt.Println(res, "sd")
}

// package main

// func main() {
// 	s := "ababcbacadefegdehijhklij"
// 	res := make(map[string]int)

// 	for _,str := {
// 		res[string(str)]++
// 	}

// 	fmt.Println(res)
// }
