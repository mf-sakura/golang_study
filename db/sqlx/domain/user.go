package domain

// User is a struct for user table.
type User struct {
	// sqlxのdbタグを使ってカラムとフィールドをマッピングする
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// Users have some users.
type Users []User
