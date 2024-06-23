package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
	if len(status) == 0 {
		return errors.New("status can not be empty")
	} else if len(status) > 140 {
		return errors.New("status exceeds 140")
	}

	return nil
}

func main() {
	err := validateStatus("This is a tweet")
	if err != nil {
		fmt.Println(err)
	}

	err = validateStatus("")
	if err != nil {
		fmt.Println(err)
	}

}
