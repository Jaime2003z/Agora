package identity

type User struct {
	ID   string
	Name string
}

func NewUser(name string) *User {
	return &User{
		ID:   generateUserID(),
		Name: name,
	}
}

func generateUserID() string {
	// Generate a more meaningful ID with username and counter
	userIDCounter++
	return "user_" + string(rune('a' + userIDCounter - 1)) // This will generate user_a, user_b, etc.
}

var userIDCounter int