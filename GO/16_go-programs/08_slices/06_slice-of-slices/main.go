package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, rows)
	for i := range matrix {
		matrix[i] = make([]int, cols)

		for j := range matrix[i] {
			matrix[i][j] = i * j
		}
	}
	return matrix
}

func test(rows, cols int) {
	fmt.Println("Output : ", createMatrix(rows, cols))
}

func main() {
	test(3, 3)
	test(4, 4)
	test(5, 3)
}
