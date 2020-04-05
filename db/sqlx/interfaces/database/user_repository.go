package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/db/sqlx/domain"
)

// Store is a function for creating a user.
func Store(db *sqlx.DB, u domain.User) (int, error) {
	// prepared statement
	// SQL Injection対策
	stmt, err := db.Prepare("INSERT INTO users (first_name, last_name) VALUES (?,?)")
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

	// // `db.Exec`でもクエリ実行は可能
	// res, err := db.Exec("INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName)
	// if err != nil {
	// 	return 0, err
	// }

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(lastInsertID), nil
}

// Update is a function for updating a user.
func Update(db *sqlx.DB, user *domain.User) error {
	// prepared statement
	// SQL Injection対策
	sql := "UPDATE `users` SET `first_name` = ?, `last_name` = ? WHERE `id` = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		return err
	}
	// 関数終了時にstatementをcloseする
	defer stmt.Close()
	// SQL文実行
	_, err = stmt.Exec(user.FirstName, user.LastName, user.ID)
	if err != nil {
		return err
	}
	return nil
}

// FindByID is a function for getting a user.
func FindByID(db *sqlx.DB, identifier int) (*domain.User, error) {
	var user domain.User
	// https://godoc.org/github.com/jmoiron/sqlx#DB.Get
	if err := db.Get(&user, "SELECT id, first_name, last_name FROM users WHERE id = ? limit 1", identifier); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
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
