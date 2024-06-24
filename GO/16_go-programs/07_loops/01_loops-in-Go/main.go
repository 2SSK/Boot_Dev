package main

import "fmt"

func bulkSend(numMessages int) float64 {
	res := 0.00

	for i := 0; i < numMessages; i++ {
		res = 1.0 + float64(i)*0.01
	}

	return res
}

func main() {
	fmt.Println(bulkSend(30))
	fmt.Println(bulkSend(1000000))
}
