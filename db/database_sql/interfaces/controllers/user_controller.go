package controllers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mf-sakura/golang_study/db/database_sql/domain"
	"github.com/mf-sakura/golang_study/db/database_sql/interfaces/database"
)

type UserController struct {
	db *sql.DB
}

func NewUserController(db *sql.DB) *UserController {
	return &UserController{db: db}
}

func (controller *UserController) Create(c echo.Context) error {
	u := domain.User{}
	c.Bind(&u)
	id, err := database.Store(controller.db, u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, id)
}

func (controller *UserController) Index(c echo.Context) error {
	users, err := database.FindAll(controller.db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

func (controller *UserController) Show(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := database.FindById(controller.db, id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, user)
}
