package main

import "fmt"

func sendSMsToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	costToCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	costToSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	totalCost := costToCustomer + costToSpouse
	return totalCost, nil
}

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2

	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}

	return costPerChar * len(message), nil
}

func main() {
	cost, err := sendSMsToCouple("Hello, how are you?", "Hi, I'm fine.")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Total cost:", cost)
}
