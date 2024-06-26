package main

import "fmt"

func getMessageCosts(messages []string) []float64 {
	costs := make([]float64, len(messages))

	for i, message := range messages {
		costs[i] = float64(len(message)) * 0.01
	}

	return costs
}

func test(messages []string) {
	costs := getMessageCosts(messages)
	fmt.Println("Costs : ", costs)
}

func main() {
	test([]string{"Hello", "World", "!"})
	test([]string{"Hello", "World", ""})
	test([]string{"Hello", "", "!"})
}
