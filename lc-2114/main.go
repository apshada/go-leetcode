// Golang program to illustrate the
// string matching using regexp
// in-built functions
package main

import (
	"fmt"
	"regexp"
)

func main() {
	// returns a copy and replaces
	// matches with the replacement string

	sen := [4]string{
		"alice and bob love leetcode",
		"this is great thanks very much",
		"i think so too",
	}
	maxLen := 0
	for i := 0; i < len(sen)-1; i++ {
		re3, _ := regexp.Compile(" ")
		match5 := re3.ReplaceAllString(sen[i], "")
		fmt.Println(match5, len(match5))
		if len(match5) > maxLen {
			maxLen = len(match5)
		}
	}

	fmt.Println(maxLen)

}
