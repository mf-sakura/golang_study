package database

import (
	"database/sql"

	"github.com/mf-sakura/golang_study/db/database_sql/domain"
)

type UserRepository struct {
	db *sql.DB
}

func Store(db *sql.DB, u domain.User) (id int, err error) {
	result, err := db.Exec(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	return
}

func FindById(db *sql.DB, identifier int) (user domain.User, err error) {
	row, err := db.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return user, err
	}
	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return user, err
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return user, nil
}

func FindAll(db *sql.DB) (users domain.Users, err error) {
	rows, err := db.Query("SELECT id, first_name, last_name FROM users")
	defer rows.Close()
	if err != nil {
		return domain.Users{}, err
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := domain.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return users, nil
}
