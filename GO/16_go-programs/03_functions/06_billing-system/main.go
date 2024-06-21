package main

import "fmt"

func calcualteFinalBill(costPerMessage float64, numMessages int) float64 {
	baseBill := calculateBaseBill(costPerMessage, numMessages)
	discount := calculateDiscount(numMessages)
	return baseBill * (1 - discount)
}

func calculateDiscount(messagesSent int) float64 {
	if messagesSent >= 5000 {
		return 0.2
	}
	if messagesSent >= 1000 {
		return 0.1
	}
	return 0

}

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}

func main() {
	test(0.2, 100)
	test(0.2, 1000)
	test(0.2, 5000)
	test(0.2, 2000)
}

func test(costPerMessage float64, numMessages int) {
	fmt.Printf("\nCost per message: %.2f\n", costPerMessage)
	fmt.Printf("Number of messages: %d\n", numMessages)
	fmt.Printf("Base bill: %.2f\n", calculateBaseBill(costPerMessage, numMessages))
	fmt.Printf("Discount: %.2f\n", calculateDiscount(numMessages))
	fmt.Printf("Final bill: %.2f\n", calcualteFinalBill(costPerMessage, numMessages))
}
