package controllers

import (
	"strconv"
	"net/http"
	"database/sql"
	
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"

	"github.com/mf-sakura/golang_study/db/api/interfaces/database"
	"github.com/mf-sakura/golang_study/db/api/domain"
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
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	id, err := database.Store(controller.db, user)
	if err != nil {
		return err
	}
	user.ID = id
	return c.JSON(http.StatusOK, user)
}

func (controller *UserController) Update(c echo.Context) (err error) {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
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

	return c.JSON(http.StatusOK, user)
}

// Index is a function for returning all users.
func (controller *UserController) Index(c echo.Context) error {
	users, err := database.FindAll(controller.db)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}

// Show is a function for returning a user.
func (controller *UserController) Show(c echo.Context) error {
	id := c.Param("id")
	identifier, err := strconv.Atoi(id)
	user, err := database.FindByID(controller.db, identifier)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	if err = c.Bind(user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}
	return c.JSON(http.StatusOK, user)
}
