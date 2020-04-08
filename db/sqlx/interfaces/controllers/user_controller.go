package controllers

import (
	"strconv"
	"database/sql"

	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/db/sqlx/domain"
	"github.com/mf-sakura/golang_study/db/sqlx/interfaces/database"
)

// UserController is a struct for db connection
type UserController struct {
	db *sqlx.DB
}

// NewUserController create a struct , UserController
func NewUserController(db *sqlx.DB) *UserController {
	return &UserController{db: db}
}

// Create is a function for creating a user.
func (controller *UserController) Create(firstName string, lastName string) (*domain.User, error) {
	user := domain.User{
		FirstName: firstName,
		LastName:  lastName,
	}
	id, err := database.Store(controller.db, user)
	if err != nil {
		return nil, err
	}
	user.ID = id
	return &user, nil
}

func (controller *UserController) Update(id string, firstName string, lastName string) (err error) {
	identifier, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	user := domain.User{
		ID: identifier,
		FirstName: firstName,
		LastName:  lastName,
	}
	tx := controller.db.MustBegin()
	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			if rollbackErr != sql.ErrTxDone {
				err = rollbackErr
			}
		}
	}()

	err = database.Update(tx, user)
	if err != nil {
		return
	}
	tx.Commit()

	return
}

// Index is a function for returning all users.
func (controller *UserController) Index() ([]domain.User, error) {
	users, err := database.FindAll(controller.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Show is a function for returning a user.
func (controller *UserController) Show(id string) (*domain.User, error) {
	// idをintegerにcastする
	identifier, err := strconv.Atoi(id)
	user, err := database.FindByID(controller.db, identifier)
	if err != nil {
		return nil, err
	}
	return user, nil
}
