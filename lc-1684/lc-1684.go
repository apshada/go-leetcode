package main

func main() {
	allowed := "ab"
	words := []string{"ad", "bd", "aaab", "baa", "badab"}

	allowedMap := make(map[string]int, 0)
	for i := 0; i < len(allowed); i++ {
		allowedMap[string(allowed[i])]++
	}

	for i := 0; i < len(words); i++ {
		// checkWords := make(map[string]int, 0)
		for j := 0; j < len(words[i]); j++ {
			// checkWords[string(checkWords[i][j])]++
		}
	}
}
