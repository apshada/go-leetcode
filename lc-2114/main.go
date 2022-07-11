// Golang program to illustrate the
// string matching using regexp
// in-built functions
package main

import (
	"fmt"
)

func main() {
	// returns a copy and replaces
	// matches with the replacement string

	sentences := []string{
		"alice and bob love leetcode", "i think so too", "this is great thanks very much",
	}

	res := mostWordsFound(sentences)
	fmt.Println(res)
}

func mostWordsFound(sentences []string) int {
	maxLen := 0
	for i := 0; i < len(sentences); i++ {
		count := 0
		for j := 0; j < len(sentences[i]); j++ {
			if string(sentences[i][j]) == " " {
				count++
			}
		}
		if count > maxLen {
			maxLen = count
		}
	}
	return maxLen + 1
}
