package domain

// User is a struct for user table.
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// Users have some users.
type Users []User
