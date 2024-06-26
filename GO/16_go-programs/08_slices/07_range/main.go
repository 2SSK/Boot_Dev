package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
	contains := func(slice []string, word string) bool {
		for _, v := range slice {
			if v == word {
				return true
			}
		}
		return false
	}

	for i, w := range msg {
		if contains(badWords, w) {
			return i
		}
	}
	return -1
}

func test(msg []string, badWords []string) {
	fmt.Println("Output : ", indexOfFirstBadWord(msg, badWords))
}

func main() {
	test([]string{"hello", "world", "this", "is", "a", "test"}, []string{"is"})
	test([]string{"hello", "world", "this", "is", "a", "test"}, []string{"world"})
}
