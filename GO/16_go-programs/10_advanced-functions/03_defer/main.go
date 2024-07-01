package main

const (
	logDelete   = "user deleted"
	logNotFound = "user not found"
	logAdmin    = "Admin deleted"
)

type user struct {
	name   string
	number int
	admin  bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)

	user, ok := users[name]
	if !ok {
		return logNotFound
	}

	if user.admin {
		return logAdmin
	}

	return logDelete
}

func test(users map[string]user, names []string) {
	for _, name := range names {
		log := logAndDelete(users, name)
		println(log)
	}
}

func main() {
	test(map[string]user{"user1": user{"user1", 1, false}, "user2": user{"user2", 2, true}}, []string{"user1", "user2", "user3"})
	test(map[string]user{"user3": user{"user3", 3, false}, "user4": user{"user4", 4, true}}, []string{"user3", "user4", "user5"})
}
