package main

import "fmt"

func getSMSErrorString(cost float64, recipient string) string {
	err := fmt.Sprintf("SMS that costs $%v to be sent to '%v' can not be sent", cost, recipient)

	return err
}

func main() {
	fmt.Println(getSMSErrorString(0.5, "1234567890"))
}
