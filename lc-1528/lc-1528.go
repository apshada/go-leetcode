package main

import (
	"fmt"
	"sort"
)

func main() {
	s := "codeleet"
	indices := []int{4, 5, 6, 7, 0, 2, 1, 3}

	restoreString(s, indices)
}

func restoreString(s string, indices []int) string {
	resMap := make(map[int]string, 0)
	str := ""
	for i, index := range indices {
		resMap[index] = string(s[i])
	}
	// fmt.Println(resMap)
	sort.Ints(indices)
	for _, val := range indices {
		// fmt.Println(resMap[val])
		str = str + resMap[val]
		// str.join()
	}
	fmt.Println(str)
	return str
}
