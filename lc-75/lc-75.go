package main

import "fmt"

func main() {
	// nums := []int{2, 0, 2, 1, 1, 0}
	nums := []int{2, 0, 1}
	low, high, mid := 0, len(nums)-1, 0

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else {
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		}
	}
	fmt.Println(nums)
}

func sortColors(nums []int) {
	low := 0
	mid := 0
	high := len(nums) - 1
	for mid <= high {
		if nums[mid] == 0 {
			swap(&nums[low], &nums[mid])
			low++
			mid++
		} else if nums[mid] == 1 {
			mid++
		} else if nums[mid] == 2 {
			swap(&nums[mid], &nums[high])
			high--
		}
	}
}

func swap(a, b *int) {
	x := *a
	*a = *b
	*b = x
}

// [2, 0, 1]
// [1, 0, 2]
// [0, 1, 2]
