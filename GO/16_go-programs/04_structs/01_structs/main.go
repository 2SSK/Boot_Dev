package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func main() {
	message := messageToSend{}

	message.phoneNumber = 6262172262
	message.message = "Ka Ho!"

	fmt.Print(message, "\n")
}
