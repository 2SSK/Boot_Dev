package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func test(numbers []int) {
	fmt.Println("Input:")
	for _, number := range numbers {
		println(number)
	}

	fmt.Println("Sum:")
	add := adder()
	for _, number := range numbers {
		println(add(number))
	}
}

func main() {
	test([]int{1, 2, 3})
	test([]int{4, 5, 6})
}
