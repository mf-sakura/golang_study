package database

import (
	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/db/sqlx/domain"
)

func Store(db *sqlx.DB, u domain.User) (id int, err error) {
	stmt, err := db.Prepare("INSERT INTO users (first_name, last_name) VALUES (?,?)")
	if err != nil {
		return id, err
	}
	defer stmt.Close()
	stmt.Exec(u.FirstName, u.LastName)
	return u.ID, nil
}

func FindById(db *sqlx.DB, identifier int) (domain.User, error) {
	var user domain.User
	if err := db.Get(&user, "SELECT id, first_name, last_name FROM users WHERE id = ? limit 1", identifier); err != nil {
		return user, err
	}
	return user, nil
}

func FindAll(db *sqlx.DB) (domain.Users, error) {
	var users []domain.User
	if err := db.Select(&users, "SELECT id, first_name, last_name FROM users"); err != nil {
		return users, err
	}
	return users, nil
}
