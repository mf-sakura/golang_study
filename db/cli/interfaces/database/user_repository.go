package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/db/cli/domain"
)

// FindAll is a function for getting all users.
func FindAll(db *sqlx.DB) (domain.Users, error) {
	var users []domain.User
	// https://godoc.org/github.com/jmoiron/sqlx#DB.Select
	if err := db.Select(&users, "SELECT id, first_name, last_name FROM users"); err != nil {
		return nil, err
	}
	return users, nil
}

// Update is a function for editing all users.
func Update(db *sqlx.DB, u domain.User) error {
	_, err := db.Exec("UPDATE users SET first_name = ?, last_name = ? WHERE id = ?", u.FirstName, u.LastName, u.ID)
	if err != nil {
		return err
	}
	return nil
}
