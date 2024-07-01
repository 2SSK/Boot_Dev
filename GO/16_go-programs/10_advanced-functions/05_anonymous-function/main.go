package main

import "fmt"

func printReports(message []string) {
	for _, message := range message {
		printCostreport(func(msg string) int {
			return 2 * len(msg)
		}, message)
	}
}

func test(messages []string) {
	defer fmt.Println("================================")
	printReports(messages)
}

func main() {
	test([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"There's no place like home",
	})

	test([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die!",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day",
	})
}

func printCostreport(costCalculator func(string) int, message string) {
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()

}
