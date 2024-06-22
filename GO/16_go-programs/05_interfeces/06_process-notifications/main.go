package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (d directMessage) importance() int {
	if d.isUrgent {
		return 50
	}

	return d.priorityLevel
}

func (g groupMessage) importance() int {
	return g.priorityLevel
}

func (s systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch t := n.(type) {
	case directMessage:
		return t.senderUsername, t.importance()
	case groupMessage:
		return t.groupName, t.importance()
	case systemAlert:
		return t.alertCode, t.importance()
	default:
		return "", 0
	}
}

func main() {
	dm := directMessage{
		senderUsername: "saurav",
		messageContent: "Hello, World!",
		priorityLevel:  10,
		isUrgent:       true,
	}

	gm := groupMessage{
		groupName:      "golang",
		messageContent: "Hello, Gophers!",
		priorityLevel:  20,
	}

	sa := systemAlert{
		alertCode:      "SYS-001",
		messageContent: "System is down!",
	}

	fmt.Println(processNotification(dm))
	fmt.Println(processNotification(gm))
	fmt.Println(processNotification(sa))
}
