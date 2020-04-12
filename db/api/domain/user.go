package domain

// User is a struct for user table.
type User struct {
	// sqlxのdbタグを使ってカラムとフィールドをマッピングする
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"first_name"`
	LastName  string `db:"last_name" json:"last_name"`
}

// Users have some users.
type Users []User
