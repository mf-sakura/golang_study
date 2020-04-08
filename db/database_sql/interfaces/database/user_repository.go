package database

import (
	"database/sql"

	"github.com/mf-sakura/golang_study/db/database_sql/domain"
)

// Update is a function for updating a user
func Update(db *sql.DB, u domain.User) (err error) {
	_, err = db.Exec(
		"update users set first_name = ?, last_name = ? where id = ?", u.FirstName, u.LastName, u.ID,
	)
	if err != nil {
		return
	}

	return
}

// Store is a function for creating a user.
func Store(db *sql.DB, u domain.User) (int, error) {
	// `?`はmysqlのplaceholder
	// `?`に第2引数以降の値を順番に代入したクエリが実行される
	result, err := db.Exec(
		"INSERT INTO users (first_name, last_name) VALUES (?,?)", u.FirstName, u.LastName,
	)
	if err != nil {
		return 0, err
	}
	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id64), nil
}

// FindByID is a function for getting a user.
func FindByID(db *sql.DB, identifier int) (*domain.User, error) {
	row, err := db.Query("SELECT id, first_name, last_name FROM users WHERE id = ?", identifier)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var id int
	var firstName string
	var lastName string
	// 次の行へ
	if !row.Next() {
		// 0件の場合はnilを返す
		return nil, nil
	}
	// selectした値をコピー
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return nil, err
	}
	user := &domain.User{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
	}
	return user, nil
}

// FindAll is a function for getting all users.
func FindAll(db *sql.DB) (domain.Users, error) {
	rows, err := db.Query("SELECT id, first_name, last_name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// 行がなくなるまでループさせる
	// https://godoc.org/database/sql#Rows.Next
	var users domain.Users
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			return nil, err
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
