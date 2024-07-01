package main

import "fmt"

/*
NOTE:
		getLogger takes a function that formats two strings into a single string
		and returns a function that formats two strings but prints the result
		instead of returning it
*/

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(s1, s2 string) {
		result := formatter(s1, s2)
		fmt.Println(result)
	}
}

func formatter(s1, s2 string) string {
	return s1 + s2
}

func test(first string, errors []error, formatter func(string, string) string) {
	defer fmt.Println("==============================")
	logger := getLogger(formatter)
	fmt.Println("Logs:")
	for _, err := range errors {
		logger(first, err.Error())
	}
}

func main() {
	test("Error: ", []error{fmt.Errorf("error1"), fmt.Errorf("error2"), fmt.Errorf("error3")}, formatter)
	test("Error: ", []error{fmt.Errorf("error4"), fmt.Errorf("error5"), fmt.Errorf("error6")}, formatter)
}
