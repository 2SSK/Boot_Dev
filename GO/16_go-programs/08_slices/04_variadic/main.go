package main

import "fmt"

func sum(nums ...int) int {
	totalsum := 0

	for i := 0; i < len(nums); i++ {
		totalsum += nums[i]
	}

	return totalsum
}

func test(nums ...int) {
	fmt.Println("Output : ", sum(nums...))
}

func main() {
	test(1, 2, 3, 4, 5)
	test(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	test(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
}
