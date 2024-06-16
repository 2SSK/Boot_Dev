package main

import "fmt"

func main() {
	messagesFromDoris := []string{
		"You doing anything later??",
		"Did you get my last message?",
		"Don't leave me hanging...",
		"Please respond I'm lonely!",
	}

	numMessages := float64(len(messagesFromDoris))
	costPreMessage := 0.02

	totalCost := numMessages * costPreMessage

	fmt.Printf("Doris spent $%.2f on text messages today\n", totalCost)
}
