package main

import (
	"fmt"
	"strings"
)

func countDistinctWords(messages []string) int {
	wordsMap := make(map[string]struct{})

	for _, message := range messages {
		words := strings.Fields(strings.ToLower(message))
		for _, word := range words {
			wordsMap[word] = struct{}{}
		}
	}

	return len(wordsMap)
}

func test(messages []string) {
	fmt.Printf("\nInput : %v\n", messages)

	fmt.Printf("\nOutput : %v\n", countDistinctWords(messages))
}

func main() {
	test([]string{"Hello, how are you?", "I am fine. How about you?", "I am good."})
	test([]string{"Hello, how are you?", "I am fine. How about you?", "I am good.", "I am fine."})
}
