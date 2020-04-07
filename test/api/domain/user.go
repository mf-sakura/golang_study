package domain

// User is a struct for user table.
// jsonタグをつけましょう
type User struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

// Users have some users.
type Users []User
