package main

import "fmt"

// ?
func (u User) sendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	}

	return "", false
}

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             membershipType
	MessageCharLimit int
}

type membershipType string

const (
	TypeStandard membershipType = "standard"
	TypePremium  membershipType = "premium"
)

func newUser(name string, membershipType membershipType) User {
	membership := Membership{Type: membershipType}
	if membershipType == TypeStandard {
		membership.MessageCharLimit = 100
	} else if membershipType == TypePremium {
		membership.MessageCharLimit = 1000
	}
	return User{Name: name, Membership: membership}
}

func test(name, message string, memberShip membershipType) {
	user := newUser(name, memberShip)

	messageLength := len(message)
	user.sendMessage(message, messageLength)

	fmt.Printf("User: %s\n", user.Name)
	fmt.Printf("Message: %s\n", message)
	fmt.Printf("Message Length: %d\n", messageLength)
	fmt.Printf("Message Char Limit: %d\n", user.MessageCharLimit)
	fmt.Printf("Membership: %s\n", memberShip)
	fmt.Printf(user.sendMessage(message, messageLength))
	println("\n")
}

func main() {
	test("saurav", "ka ho", TypePremium)
	test("saurav", "re saar", TypeStandard)
	test("saurav", "bhak bsdwala", "")
}
