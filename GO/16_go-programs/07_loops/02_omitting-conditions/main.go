package main

import "fmt"

func maxMessages(thresh int) {
	res := 0
	for i := 0; i < thresh; i++ {
		res += 100 + i
		fmt.Printf("i: %d, res: %d\n", i, res)
	}
}

func main() {
	maxMessages(10)
}
