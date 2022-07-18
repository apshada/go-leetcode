package main

import "fmt"

func main() {
	items := [][]string{
		{"phone", "blue", "pixel"},
		{"computer", "silver", "lenovo"},
		{"phone", "gold", "iphone"},
	}
	ruleKey := "type"
	ruleValue := "phone"
	count := 0
	setSearchCol := 0
	if ruleKey == "type" {
		setSearchCol = 0
	}
	if ruleKey == "color" {
		setSearchCol = 1
	}
	if ruleKey == "name" {
		setSearchCol = 2
	}

	for i := 0; i < len(items); i++ {
		if items[i][setSearchCol] == ruleValue {
			count++
		}

	}
	fmt.Println(count)
}
