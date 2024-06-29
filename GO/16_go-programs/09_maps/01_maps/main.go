package main

import (
	"errors"
	"fmt"
)

func getUserMap(names []string, phoneNumber []int) (map[string]user, error) {
	// Check if the lengths of the slices are equal
	if len(names) != len(phoneNumber) {
		return nil, errors.New("invalid sizes")
	}

	// Initialze the map to store user structs
	userMap := make(map[string]user)

	// Iterate over the slices and populate the map
	for i := 0; i < len(names); i++ {
		userMap[names[i]] = user{
			name:        names[i],
			phoneNumber: phoneNumber[i],
		}
	}

	return userMap, nil
}

type user struct {
	name        string
	phoneNumber int
}

func test(names []string, phoneNumber []int) {
	userMap, err := getUserMap(names, phoneNumber)
	if err != nil {
		fmt.Println("Output : ", err)
		return
	}
	for _, user := range userMap {
		fmt.Println("Output : ", user)
	}
}

func main() {
	test([]string{"John", "Doe"}, []int{123, 456})
	test([]string{"John", "Doe", "Jane"}, []int{123, 456})
	test([]string{"John", "Doe", "Jane"}, []int{123, 456, 789})
}
