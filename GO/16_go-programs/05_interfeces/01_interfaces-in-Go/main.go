package main

import (
	"fmt"
	"time"
)

type message interface {
	getMessage() string
}

func sendMessage(msg message) (string, int) {
	messageContent := msg.getMessage()
	messageCost := len(messageContent) * 3
	return messageContent, messageCost
}

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func main() {
	bm := birthdayMessage{
		birthdayTime:  time.Now(),
		recipientName: "John",
	}

	sr := sendingReport{
		reportName:    "Monthly",
		numberOfSends: 150,
	}

	msgContent, msgCost := sendMessage(bm)
	fmt.Printf("Message: %s\nCost: %d\n", msgContent, msgCost)

	msgContent, msgCost = sendMessage(sr)
	fmt.Printf("Message: %s\nCost: %d\n", msgContent, msgCost)
}
