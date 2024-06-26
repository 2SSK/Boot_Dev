package main

import "fmt"

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	maxDay := 0
	for _, c := range costs {
		if c.day > maxDay {
			maxDay = c.day
		}
	}

	dailyCosts := make([]float64, maxDay+1)

	for _, c := range costs {
		dailyCosts[c.day] += c.value
	}

	return dailyCosts
}

func test(costs []cost) {
	fmt.Println("Output : ", getCostsByDay(costs))
}

func main() {
	test([]cost{{1, 1.0}, {2, 2.0}, {3, 3.0}, {4, 4.0}, {5, 5.0}})
	test([]cost{{1, 1.0}, {2, 2.0}, {3, 3.0}, {4, 4.0}, {5, 5.0}, {6, 6.0}, {7, 7.0}, {8, 8.0}, {9, 9.0}, {10, 10.0}})
	test([]cost{{1, 1.0}, {2, 2.0}, {3, 3.0}, {4, 4.0}, {5, 5.0}, {6, 6.0}, {7, 7.0}, {8, 8.0}, {9, 9.0}, {10, 10.0}, {11, 11.0}})
}
