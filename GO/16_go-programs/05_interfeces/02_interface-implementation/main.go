package main

type employee interface {
	getName() string
	getSalary() int
}

type contractor struct {
	name         string
	hourlyPay    int
	hoursPerYear int
}

func (c contractor) getName() string {
	return c.name
}

func (c contractor) getSalary() int {
	return c.hourlyPay * c.hoursPerYear
}

type fullTime struct {
	name   string
	salary int
}

func (ft fullTime) getName() string {
	return ft.name
}

func (ft fullTime) getSalary() int {
	return ft.salary
}

func main() {
	c := contractor{
		name:         "John",
		hourlyPay:    10,
		hoursPerYear: 2000,
	}

	ft := fullTime{
		name:   "Jane",
		salary: 50000,
	}

	employees := []employee{c, ft}

	for _, e := range employees {
		println(e.getName(), e.getSalary())
	}
}
