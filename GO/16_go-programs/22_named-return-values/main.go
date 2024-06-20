package main

import "fmt"

func yearsUntilEvents(age int) (int, int, int) {
	yearsUntilAdult := 18 - age
	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}

	yearsUntilDrinking := 21 - age
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}

	yearsUntilCarRental := 25 - age
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental
}

func main() {
	tests(21)
	tests(17)
	tests(26)
}

func tests(age int) {
	fmt.Println(yearsUntilEvents(age))
}
