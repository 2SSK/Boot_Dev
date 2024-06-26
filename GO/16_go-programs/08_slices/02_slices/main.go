package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	msgSlice := messages[:]

	if plan == planPro {
		return msgSlice, nil
	} else if plan == planFree {
		return msgSlice[:2], nil
	} else {
		return nil, errors.New("unsupported plan")
	}
}

func test(plan string, messages [3]string) {
	msg, err := getMessageRetriesForPlan(plan, messages)
	if err != nil {
		fmt.Println("Error : ", err)
		return
	}
	fmt.Println("Message : ", msg)
}

func main() {
	test(planPro, [3]string{"Hello", "World", "!"})
	test(planPro, [3]string{"Hello", "World", ""})
	test(planPro, [3]string{"Hello", "", "!"})
	test(planFree, [3]string{"Hello", "World", "!"})
	test(planFree, [3]string{"Hello", "World", ""})
	test(planFree, [3]string{"Hello", "", "!"})
	test("basic", [3]string{"Hello", "World", "!"})
}
