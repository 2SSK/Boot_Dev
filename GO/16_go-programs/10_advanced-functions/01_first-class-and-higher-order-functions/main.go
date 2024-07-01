package main

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}

	return formattedMessages
}

func formatter(message string) string {
	return "formatted: " + message
}

func test(messages []string) {
	formattedMessages := getFormattedMessages(messages, formatter)
	for _, message := range formattedMessages {
		println(message)
	}
}

func main() {
	test([]string{"message1", "message2", "message3"})
	test([]string{"message4", "message5", "message6"})
}
