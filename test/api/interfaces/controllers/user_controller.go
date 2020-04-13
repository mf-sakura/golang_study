package controllers

import (
	"database/sql"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/mf-sakura/golang_study/test/api/domain"
	"github.com/mf-sakura/golang_study/test/api/interfaces/database"
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
func (controller *UserController) Create(c echo.Context) (err error) {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	tx := controller.db.MustBegin()
	defer func() {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			// tx が閉じていることにより
			// rollback には失敗するが特に問題はないのでエラーは返さない
			if rollbackErr != sql.ErrTxDone {
				err = rollbackErr
			}
		}
	}()
	id, err := database.Store(tx, user)
	if err != nil {
		return
	}
	user.ID = id
	if err := tx.Commit(); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, &user)
}

// Index is a function for returning all users.
func (controller *UserController) Index(c echo.Context) error {
	firstName := c.QueryParam("first_name")
	var err error
	var users domain.Users

	if firstName != "" {
		users, err = database.FirstNameLike(controller.db, firstName)
	} else {
		users, err = database.FindAll(controller.db)
	}

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &users)
}
