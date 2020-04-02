package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/db/sqlx/domain"
)

// Store is a function for creating a user.
func Store(db *sqlx.DB, u domain.User) (id int, err error) {
	// prepared statement
	stmt, err := db.Prepare("INSERT INTO users (first_name, last_name) VALUES (?,?)")
	if err != nil {
		return id, err
	}
	// 関数終了時にstatementをcloseする
	defer stmt.Close()
	// SQL文実行
	stmt.Exec(u.FirstName, u.LastName)
	return u.ID, nil
}

// FindByID is a function for getting a user.
func FindByID(db *sqlx.DB, identifier int) (domain.User, error) {
	var user domain.User
	// https://godoc.org/github.com/jmoiron/sqlx#DB.Get
	if err := db.Get(&user, "SELECT id, first_name, last_name FROM users WHERE id = ? limit 1", identifier); err != nil {
		return user, err
	}
	return user, nil
}

// FindAll is a function for getting all users.
func FindAll(db *sqlx.DB) (domain.Users, error) {
	var users []domain.User
	// https://godoc.org/github.com/jmoiron/sqlx#DB.Select
	if err := db.Select(&users, "SELECT id, first_name, last_name FROM users"); err != nil {
		return users, err
	}
	return users, nil
}
