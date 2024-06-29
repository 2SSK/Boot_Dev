package main

import (
	"errors"
	"fmt"
)

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	u, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}

	if !u.scheduledForDeletion {
		return false, nil
	}

	delete(users, name)
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func test(users map[string]user, name string) {
	deleted, err := deleteIfNecessary(users, name)

	fmt.Printf("\nInput : %v\n", name)

	if err != nil {
		fmt.Println("Output : ", err)
		return
	}

	if deleted {
		fmt.Println("Output : ", "Deleted")
		return
	}

	fmt.Println("Output : ", "Not deleted")
}

func main() {
	users := map[string]user{
		"Mohit":  {name: "Mohit", number: 123, scheduledForDeletion: false},
		"Amit":   {name: "Amit", number: 456, scheduledForDeletion: true},
		"Saurav": {name: "Saurav", number: 789, scheduledForDeletion: false},
	}

	test(users, "Amit")
	test(users, "Ahbab")
	test(users, "Saurav")
	test(users, "Mohit")
}
