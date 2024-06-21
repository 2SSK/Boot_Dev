package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	if mToSend.sender.name == "" || mToSend.sender.number == 0 || mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}
	return true
}

func main() {
	test("Saurav", "Mohit", "bsdk", 123, 456)
	test("Saurav", "Mohit", "bsdk", 123, 0)
	test("Saurav", "Mohit", "bsdk", 0, 456)
	test("Saurav", "Mohit", "", 123, 456)
	test("Saurav", "", "bsdk", 123, 456)
	test("", "Mohit", "bsdk", 123, 456)
}

func test(nam1, nam2, msg string, no1, no2 int) {
	message := messageToSend{}

	message.sender = user{
		name:   nam1,
		number: no1,
	}

	message.recipient = user{
		name:   nam2,
		number: no2,
	}

	message.message = msg

	if canSendMessage(message) {
		fmt.Println("Message sent!")
	} else {
		fmt.Println("Message not sent!")
	}
}
