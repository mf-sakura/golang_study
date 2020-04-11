package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/test/api/domain"
)

// Store is a function for creating a user.
func Store(tx *sqlx.Tx, u domain.User) (int, error) {
	// prepared statement
	// SQL Injection対策
	stmt, err := tx.Prepare("INSERT INTO users (first_name, last_name) VALUES (?,?)")
	if err != nil {
		return 0, err
	}

	// 関数終了時にstatementをcloseする
	defer stmt.Close()
	// SQL文実行
	res, err := stmt.Exec(u.FirstName, u.LastName)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertID), nil
}

// FindAll is a function for getting all users.
func FindAll(db *sqlx.DB) (domain.Users, error) {
	var users []domain.User
	// https://godoc.org/github.com/jmoiron/sqlx#DB.Select
	if err := db.Select(&users, "SELECT id, first_name, last_name FROM users"); err != nil {
		return nil, err
	}
	return users, nil
}

// FirstNameLike is a function for getting all users whose first name is like ?.
func FirstNameLike(db *sqlx.DB, firstName string) (domain.Users, error) {
	var users []domain.User

	// https://godoc.org/github.com/jmoiron/sqlx#DB.Select
	if err := db.Select(&users, "SELECT id, first_name, last_name FROM users WHERE first_name LIKE ?", "%"+firstName+"%"); err != nil {
		return nil, err
	}
	return users, nil
}
