package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	authString := fmt.Sprintf("Authorization: Basic %s : %s\n", a.username, a.password)

	return authString
}

func test(username, password string) {
	auth := authenticationInfo{username, password}

	fmt.Println(auth.getBasicAuth())
}

func main() {
	test("user1", "password1")
	test("user2", "password2")
	test("user3", "password3")
}
