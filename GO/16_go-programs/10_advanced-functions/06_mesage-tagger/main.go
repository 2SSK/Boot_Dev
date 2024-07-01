package main

import "strings"

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, message := range messages {
		messages[i].tags = tagger(message)
	}

	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	if strings.Contains(strings.ToLower(msg.content), "urgent") {
		tags = append(tags, "Urgent")
	}

	if strings.Contains(strings.ToLower(msg.content), "sale") {
		tags = append(tags, "Promo")
	}

	return tags
}

func test(messages []sms) {
	taggedMessages := tagMessages(messages, tagger)
	for _, message := range taggedMessages {
		println(message.id, message.content, message.tags)
	}
}

func main() {
	test([]sms{42: sms{"1", "Urgent: Sale today!", nil}, 43: sms{"2", "Sale today!", nil}, 44: sms{"3", "No sale today!", nil}})

	test([]sms{45: sms{"4", "Urgent: Sale today!", nil}, 46: sms{"5", "Sale today!", nil}, 47: sms{"6", "No sale today!", nil}})
}
