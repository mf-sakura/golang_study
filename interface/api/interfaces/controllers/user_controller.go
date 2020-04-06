package controllers

import (
	"errors"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/mf-sakura/golang_study/interface/api/domain"
	"github.com/mf-sakura/golang_study/interface/api/interfaces/database"
)

// UserController is a struct for db connection
type UserController struct {
	db         *sqlx.DB
	repository database.UserRepository
}

// NewUserController create a struct , UserController
func NewUserController(db *sqlx.DB, repository database.UserRepository) *UserController {
	return &UserController{
		db:         db,
		repository: repository,
	}
}

// Create is a function for creating a user.
func (controller *UserController) Create(c echo.Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	if user.LastName == "" || user.FirstName == "" {
		return echo.NewHTTPError(http.StatusBadRequest, errors.New("Both Name must be non-empty"))
	}
	id, err := controller.repository.Store(controller.db, user)
	if err != nil {
		return err
	}
	user.ID = id
	return c.JSON(http.StatusCreated, &user)
}

// Index is a function for returning all users.
func (controller *UserController) Index(c echo.Context) error {
	users, err := controller.repository.FindAll(controller.db)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &users)
}
