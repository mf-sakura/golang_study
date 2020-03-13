package controllers

import (
	"database/sql"
	"strconv"

	"github.com/mf-sakura/golang_study/db/database_sql/domain"
	"github.com/mf-sakura/golang_study/db/database_sql/interfaces/database"
)

type UserController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{db: db}
}

func (controller *UserController) Create(firstName string, lastName string) (domain.User, error) {
	user := domain.User{
		FirstName: firstName,
		LastName:  lastName,
	}
	id, err := database.Store(controller.db, user)
	if err != nil {
		return domain.User{}, err
	}
	user.ID = id
	return user, nil
}

func (controller *UserController) Index() ([]domain.User, error) {
	users, err := database.FindAll(controller.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (controller *UserController) Show(id string) (domain.User, error) {
	identifier, err := strconv.Atoi(id)
	user, err := database.FindById(controller.db, identifier)
	if err != nil {
		return user, err
	}
	return user, nil
}
