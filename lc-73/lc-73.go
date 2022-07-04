package main

import "fmt"

type rowCol struct {
	row int
	col int
}

func main() {
	nums := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(nums)
}

func setZeroes(matrix [][]int) {
	resMap := make(map[int]rowCol, 0)
	v := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				resMap[v] = rowCol{i, j}
				v++
			}
		}
	}

	for _, rowCol := range resMap {
		manipulateArray(rowCol.row, rowCol.col, matrix)
		fmt.Println(resMap)
	}
}

func manipulateArray(row int, column int, matrix [][]int) {
	for i := 0; i < len(matrix); i++ {
		matrix[i][column] = 0
	}
	for j := 0; j < len(matrix[row]); j++ {
		matrix[row][j] = 0
	}
	fmt.Print(matrix)

}
