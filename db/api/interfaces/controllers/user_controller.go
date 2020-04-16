package controllers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/mf-sakura/golang_study/db/api/domain"
	"github.com/mf-sakura/golang_study/db/api/interfaces/database"
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
func (controller *UserController) Create(c echo.Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	id, err := database.Store(controller.db, user)
	if err != nil {
		return err
	}
	user.ID = id
	return c.JSON(http.StatusCreated, &user)
}

// Index is a function for returning all users.
func (controller *UserController) Index(c echo.Context) error {
	users, err := database.FindAll(controller.db)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &users)
}
