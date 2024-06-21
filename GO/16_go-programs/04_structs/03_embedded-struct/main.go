package main

import "fmt"

type sender struct {
	user
	rateLimit int
}

type user struct {
	name   string
	number int
}

func main() {
	test("user1", 1, 10)
	test("user2", 123, 1)
	test("user3", 123, 0)
}

func test(name string, number, rateLimit int) {
	sender := sender{}

	sender.name = name
	sender.number = number

	sender.rateLimit = rateLimit

	fmt.Printf("%s, %d\n", sender.name, sender.number)
}
