package main

import "fmt"

func main() {
	jewels := "z"
	stones := "ZZ"
	totalJewles := 0
	resMap := make(map[string]int, 0)

	for _, stone := range stones {
		resMap[string(stone)] = resMap[string(stone)] + 1
	}

	for _, jewel := range jewels {
		totalJewles = totalJewles + resMap[string(jewel)]
	}

	fmt.Println(totalJewles)
}
