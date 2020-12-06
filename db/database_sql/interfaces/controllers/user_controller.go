package controllers

import (
	"database/sql"
	"strconv"

	"github.com/mf-sakura/golang_study/db/database_sql/domain"
	"github.com/mf-sakura/golang_study/db/database_sql/interfaces/database"
)

// UserController is a struct for db connection.
type UserController struct {
	db *sql.DB
}

// NewUserController create a struct , UserController.
func NewUserController(db *sql.DB) *UserController {
	return &UserController{db: db}
}

// Update is a function for updating a user.
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

	err = database.Update(controller.db, user)
	if err != nil {
		return
	}

	return
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
