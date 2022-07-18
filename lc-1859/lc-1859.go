package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "is2 sentence4 This1 a3"
	resStr := ""
	res := strings.Split(s, " ")
	data := make(map[int]string, 0)
	for _, r := range res {
		e := int(r[len(r)-1])
		data[e-48] = r
	}

	for i, _ := range res {
		i = i + 1
		str := data[i]
		extractStr := str[:len(str)-1]
		resStr = resStr + extractStr
		if i != len(res) {
			resStr = resStr + " "
		}
		// inputFmt:=data[:len(data)-1]
		// str = da
	}
	fmt.Println(resStr)
}
