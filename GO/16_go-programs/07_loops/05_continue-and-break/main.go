package main

import "fmt"

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func printPrimes(max int) {
	for i := 1; i <= max; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}

func test(max int) {
	fmt.Printf("Primes up to %v: \n", max)
	printPrimes(max)
	fmt.Println("======================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
