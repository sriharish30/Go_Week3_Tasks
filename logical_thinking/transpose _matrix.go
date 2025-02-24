package main

import "fmt"

func transposeMatrix(matrix [3][3]int) [3][3]int {
	var transposed [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

func main() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	result := transposeMatrix(matrix)

	for _, row := range result {
		fmt.Println(row)
	}
}
