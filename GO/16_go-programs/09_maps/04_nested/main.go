package main

import "fmt"

func getNameCounts(names []string) map[rune]map[string]int {
	// Initialize the nested map
	nameCounts := make(map[rune]map[string]int)

	// Itrerate over each name in the slice
	for _, name := range names {
		if len(name) == 0 {
			continue // Skip empty names
		}
		firstChar := rune(name[0])

		// Initialize the inner map if it does not exist
		if _, exists := nameCounts[firstChar]; !exists {
			nameCounts[firstChar] = make(map[string]int)
		}

		// Increment the count for the name
		nameCounts[firstChar][name]++
	}

	return nameCounts
}

func test(names []string) {
	fmt.Println("Before: ", names)

	nameCounts := getNameCounts(names)

	fmt.Println("After: ", nameCounts)
}

func main() {
	test([]string{"John", "Jane", "Joe", "Jack"})
	test([]string{"Mahesh", "Suresh", "Ramesh", "Jack", "Jill"})
	test([]string{"John", "Jane", "Joe", "Jack", "Jill", "Jen"})
}
