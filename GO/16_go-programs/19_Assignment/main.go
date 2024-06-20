package main

import "fmt"

func getMothlyPrice(tier string) int {
	var price int = 0

	// Check tier and determine price
	if tier == "basic" {
		price = 100
	} else if tier == "premium" {
		price = 150
	} else if tier == "enterprise" {
		price = 500
	}

	// Convert dollar into pennies
	price *= 100

	return price
}

func main() {
	test("basic")
	test("premium")
	test("enterprise")
	test("false case")
}

func test(tier string) {
	fmt.Println(getMothlyPrice(tier))
}
