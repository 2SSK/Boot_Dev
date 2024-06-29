package main

import "fmt"

func findSuggestedFriends(username string, friendships map[string][]string) []string {
	if _, exists := friendships[username]; !exists {
		return nil
	}

	directFriends := friendships[username]
	suggestedFriendsMap := make(map[string]struct{})
	directFriendsMap := make(map[string]struct{})

	for _, friend := range directFriends {
		directFriendsMap[friend] = struct{}{}
	}

	for _, friend := range directFriends {
		if _, exists := friendships[friend]; !exists {
			for _, friendOfFriend := range friendships[friend] {
				if friendOfFriend != username {
					if _, isDirectFriend := directFriendsMap[friendOfFriend]; !isDirectFriend {
						suggestedFriendsMap[friendOfFriend] = struct{}{}
					}
				}
			}
		}
	}

	sugesstedFriends := make([]string, 0, len(suggestedFriendsMap))
	for friend := range suggestedFriendsMap {
		sugesstedFriends = append(sugesstedFriends, friend)
	}

	return sugesstedFriends
}

func test(username string, friendships map[string][]string) {
	fmt.Println("\nUsername: ", username)
	fmt.Println("Friendships: ", friendships)

	fmt.Println("Suggested Friends: ", findSuggestedFriends(username, friendships))
}

func main() {
	test("John", map[string][]string{
		"Alice":   {"Bob", "Charlie"},
		"Bob":     {"Alice", "Charlie", "David"},
		"Charlie": {"Alice", "Bob", "David", "Eve"},
		"David":   {"Bob", "Charlie"},
		"Eve":     {"Charlie"},
	})
}
