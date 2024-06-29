package main

import "fmt"

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, user := range messagedUsers {
		if _, ok := validUsers[user]; ok {
			validUsers[user]++
		}
	}
}

func test(msgUsers []string, validUsers map[string]int) {
	fmt.Println("Before: ", validUsers)
	getCounts(msgUsers, validUsers)
	fmt.Println("After: ", validUsers)
}

func main() {
	validUsers := map[string]int{
		"John": 0,
		"Jane": 0,
		"Joe":  0,
		"Jack": 0,
	}

	msgUsers := []string{"John", "Jane", "Joe", "Jack"}

	test(msgUsers, validUsers)
}
