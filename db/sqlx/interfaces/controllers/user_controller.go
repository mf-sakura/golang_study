package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo/v4"
	"github.com/mf-sakura/golang_study/db/sqlx/domain"
	"github.com/mf-sakura/golang_study/db/sqlx/interfaces/database"
)

// UserController is a struct for db connection
type UserController struct {
	db *sqlx.DB
}

type userRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type userResponse struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
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

// CreateByEcho is a function for creating a user by echo.
func (controller *UserController) CreateByEcho(c echo.Context) error {
	userRequest := userRequest{}
	if err := c.Bind(&userRequest); err != nil {
		return err
	}

	if userRequest.FirstName == "" || userRequest.LastName == "" {
		return fmt.Errorf("invalid request %v", userRequest)
	}

	user := domain.User{
		FirstName: userRequest.FirstName,
		LastName:  userRequest.LastName,
	}
	id, err := database.Store(controller.db, user)
	if err != nil {
		return err
	}
	user.ID = id

	res := &userResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
	// fmt.Printf("res: %v\n", res)
	c.JSON(http.StatusOK, res)
	return nil
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

// Update is a function for updating a user.
func (controller *UserController) Update(id, firstName, lastName string) (*domain.User, error) {
	// idをintegerにcastする
	identifier, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// ここの処理いらんかも
	user, err := database.FindByID(controller.db, identifier)
	if err != nil {
		return nil, err
	}

	user.FirstName = firstName
	user.LastName = lastName
	err = database.Update(controller.db, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
