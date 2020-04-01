package controllers

import (
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/db/cli/domain"
	"github.com/mf-sakura/golang_study/db/cli/interfaces/database"
)

// UserController is a struct for db connection
type UserController struct {
	db *sqlx.DB
}

// NewUserController create a struct , UserController
func NewUserController(db *sqlx.DB) *UserController {
	return &UserController{db: db}
}

// Index is a function for returning all users.
func (controller *UserController) Index() ([]domain.User, error) {
	users, err := database.FindAll(controller.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Update is a function for editing a user.
func (controller *UserController) Update(id string, firstName string, lastName string) error {
	identifier, err := strconv.Atoi(id)
	user := domain.User{
		ID:        identifier,
		FirstName: firstName,
		LastName:  lastName,
	}
	err = database.Update(controller.db, user)
	if err != nil {
		return err
	}
	return nil
}
