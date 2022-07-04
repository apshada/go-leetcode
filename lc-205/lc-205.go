package main

import "fmt"

func main() {
	s := "egg"
	t := "add"
	res := isIsomorphic(s, t)
	fmt.Println(res)
}

func isIsomorphic(s string, t string) bool {
	map_s := make(map[int]int)
	map_t := make(map[int]int)

	for i := range s {
		if map_s[int(s[i])] != map_t[int(t[i])] {
			return false
		}
		map_s[int(s[i])] = i + 1
		map_t[int(t[i])] = i + 1
	}
	return true
}
