package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	msg := [3]string{primary, secondary, tertiary}
	var cost [3]int

	cost[0] = len(msg[0])

	for i := 1; i < 3; i++ {
		cost[i] = cost[i-1] + len(msg[i])
	}

	return msg, cost
}

func test(a, b, c string) {
	msg, cost := getMessageWithRetries(a, b, c)
	fmt.Println("Message : ", msg)
	fmt.Println("Cost : ", cost)
}

func main() {
	test("Hello", "World", "!")
	test("Hello", "World", "")
	test("Hello", "", "!")
}
