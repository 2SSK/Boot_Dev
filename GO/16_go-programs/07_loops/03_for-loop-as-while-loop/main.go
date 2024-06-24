package main

import "fmt"

func getMaxMessagesToSend(costMultiplier float64, maxCostInPennies int) int {
	actualCostInPennies := 1.0
	maxMessagesToSend := 1

	for actualCostInPennies <= float64(maxCostInPennies) {
		actualCostInPennies *= costMultiplier
		maxMessagesToSend++
	}

	if actualCostInPennies >= float64(maxCostInPennies) {
		maxMessagesToSend--
	}

	return maxMessagesToSend
}

func main() {
	fmt.Println(getMaxMessagesToSend(1.01, 100))
	fmt.Println(getMaxMessagesToSend(1.01, 1000000))
}
