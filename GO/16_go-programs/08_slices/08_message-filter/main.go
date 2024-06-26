package main

type Message interface {
	Type() string
}

type TextMessage struct {
	Sender  string
	Content string
}

func (tm TextMessage) Type() string {
	return "text"
}

type MediaMessage struct {
	Sender    string
	MediaType string
	Content   string
}

func (mm MediaMessage) Type() string {
	return "media"
}

type LinkMessage struct {
	Sender  string
	URL     string
	Content string
}

func (lm LinkMessage) Type() string {
	return "link"
}

func filterMessages(messages []Message, filterType string) []Message {
	var filteredMessages []Message

	for _, msg := range messages {
		if msg.Type() == filterType {
			filteredMessages = append(filteredMessages, msg)
		}
	}

	return filteredMessages
}

func test(messages []Message, filterType string) {
	filteredMessages := filterMessages(messages, filterType)
	for _, msg := range filteredMessages {
		switch msg.Type() {
		case "text":
			tm := msg.(TextMessage)
			println("TextMessage: ", tm.Sender, tm.Content)
		case "media":
			mm := msg.(MediaMessage)
			println("MediaMessage: ", mm.Sender, mm.MediaType, mm.Content)
		case "link":
			lm := msg.(LinkMessage)
			println("LinkMessage: ", lm.Sender, lm.URL, lm.Content)
		}
	}
}

func main() {
	test([]Message{TextMessage{"Alice", "Hello"}, MediaMessage{"Bob", "image", "image.jpg"}, LinkMessage{"Charlie", "https://example.com", "Check this out"}}, "text")
}
