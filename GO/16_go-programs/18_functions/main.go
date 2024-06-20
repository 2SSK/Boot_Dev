package main

import "fmt"

func concat(x string, y string) string {
	return x + y
}

func main() {
	test("Lane,", " happy birthday!")
	test("Elon,", " hope that Tesla thing works out")
	test("Go,", " is fantastic")
}

func test(x string, y string) {
	fmt.Println(concat(x, y))
}
