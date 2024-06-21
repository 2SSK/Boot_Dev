package main

import "fmt"

func main() {
	// intialize variables here
	var smsSendingLimit int
	var costPerSMS float32
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
