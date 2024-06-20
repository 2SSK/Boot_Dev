package main

import "fmt"

func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	lastMonthBill := getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill := getBillForMonth(costPerSend, numThisMonth)

	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messageSent int) int {
	return costPerSend * messageSent
}

func main() {
	test(10, 100, 220)
	test(10, 100, 98)
	test(10, 100, 500)
}

// Unit Test function
func test(costPerSend, numLastMonth, numThisMonth int) {
	fmt.Println(monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth))
}
