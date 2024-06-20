package main

import "fmt"

func placeOrder(productID string, quantity int, accountBalance float64) (bool, float64) {
	if quantity > amountInStock(productID) {
		return false, accountBalance
	}

	totalPrice := calcPrice(productID, quantity)

	if accountBalance < totalPrice {
		return false, accountBalance
	}

	return true, accountBalance - totalPrice
}

func calcPrice(productID string, quantity int) float64 {
	return priceList(productID) * float64(quantity)
}

func main() {
	test("1", 10, 15.00)
	test("2", 10, 15.00)
	test("3", 10, 15.00)
	test("4", 10, 15.00)
	test("5", 10, 15.00)
}

func test(productID string, quantity int, accountBalance float64) {
	fmt.Printf("\nProduct ID: %s\n", productID)
	fmt.Printf("Quantity: %d\n", quantity)
	fmt.Printf("Account balance: %.2f\n", accountBalance)
	fmt.Printf("Price per item: %.2f\n", priceList(productID))
	fmt.Printf("Amount in stock: %d\n", amountInStock(productID))
	fmt.Printf("Total price: %.2f\n", calcPrice(productID, quantity))
	success, newBalance := placeOrder(productID, quantity, accountBalance)
	fmt.Printf("Order placed: %v\n", success)
	fmt.Printf("New account balance: %.2f\n", newBalance)
}

func priceList(productID string) float64 {
	if productID == "1" {
		return 1.50
	} else if productID == "2" {
		return 2.25
	} else if productID == "3" {
		return 3.00
	} else if productID == "4" {
		return 1.00
	} else if productID == "5" {
		return 2.50
	} else if productID == "6" {
		return 8.99
	} else if productID == "7" {
		return 22.50
	} else if productID == "8" {
		return 50.00
	} else if productID == "9" {
		return 999.99
	} else {
		return 0.00
	}
}

func amountInStock(productID string) int {
	if productID == "1" {
		return 11
	} else if productID == "2" {
		return 25
	} else if productID == "3" {
		return 4
	} else if productID == "4" {
		return 6
	} else if productID == "5" {
		return 50
	} else if productID == "6" {
		return 2
	} else if productID == "7" {
		return 0
	} else if productID == "8" {
		return 99
	} else if productID == "9" {
		return 1
	} else {
		return 0
	}
}
