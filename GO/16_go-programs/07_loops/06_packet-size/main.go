package main

func getPacketSize(message string) int {
	n := len(message)

	if isPrime(n) {
		return 0
	}

	for packetSize := n; packetSize >= 1; packetSize-- {
		if n%packetSize == 0 {
			numPackets := n / packetSize
			if !isPrime(numPackets) && numPackets != 1 {
				return packetSize
			}
		}
	}

	return 0
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func test(message string) {
	packetSize := getPacketSize(message)
	if packetSize == 0 {
		println("Message cannot be sent securely")
	} else {
		println("Packet size: ", packetSize)
	}
}

func main() {
	test("Hello, world!")
	test("Hi there!")
}
