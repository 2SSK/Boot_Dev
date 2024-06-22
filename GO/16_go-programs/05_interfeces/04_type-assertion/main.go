package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	switch t := e.(type) {
	case email:
		return t.toAddress, t.cost()
	case sms:
		return t.toPhoneNumber, t.cost()
	default:
		return "", 0.0
	}

	//if c, ok := e.(email); ok {
	//return c.toAddress, c.cost()
	//} else if s, ok := e.(sms); ok {
	//return s.toPhoneNumber, s.cost()
	//
	//}
	//
	//return "", 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * 0.05
	}

	return float64(len(e.body)) * 0.01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * 0.1
	}

	return float64(len(s.body)) * 0.03
}

func (i invalid) cost() float64 {
	return 0.0
}

func main() {
	test(true, true, "Hello", "magambo@khush")
	test(true, false, "Hello", "magambo@dukh")
	test(false, true, "ka ho", "123456787")
	test(false, false, "ka ho", "12345")
}
func test(isEmail, isSub bool, body, add string) {
	ok := isEmail
	if ok {
		e := email{isSubscribed: isSub, body: body, toAddress: add}
		add, cost := getExpenseReport(e)
		fmt.Printf("Email to: %s, cost: %.2f\n", add, cost)
	} else {
		s := sms{isSubscribed: isSub, body: body, toPhoneNumber: add}
		add, cost := getExpenseReport(s)
		fmt.Printf("SMS to: %s, cost: %.2f\n", add, cost)
	}
}
