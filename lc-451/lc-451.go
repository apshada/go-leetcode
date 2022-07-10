package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := "Aabb"
	resMap := make(map[int]int, 0)
	keys := make([]int, 0)
	str := ""
	for _, str := range s {
		resMap[int(str-'0')]++
	}

	for key := range resMap {
		keys = append(keys, key)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))

	for _, key := range keys {
		c := resMap[key]
		// for c > 0 {
		// 	str = str + key
		// 	c--
		// }
		z := strconv.Itoa(key)

		// z := fmt.Sprintf("%v", key)
		fmt.Println(str, c, key, resMap, z)
	}
}
