package domain

type User struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

type Users []User
