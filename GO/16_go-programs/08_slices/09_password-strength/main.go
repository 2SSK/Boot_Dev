package main

import (
	"fmt"
	"unicode"
)

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	hasUpper := false
	hasDigit := false

	for _, char := range password {
		if unicode.IsUpper(char) {
			hasUpper = true
		}

		if unicode.IsDigit(char) {
			hasDigit = true
		}

		if hasUpper && hasDigit {
			return true
		}
	}

	return false
}

func test(password string) {
	fmt.Printf("\nPassword : %v", password)

	validitiy := isValidPassword(password)

	if validitiy == true {
		fmt.Println("\n\tPassword is valid")
	} else {
		fmt.Println("\n\tPassword is invalid")
	}
}

func main() {
	test("ssk")
	test("6206860109")
	test("Saurav@12345")
	test("saurav#12")
}
