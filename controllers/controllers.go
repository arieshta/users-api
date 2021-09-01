package controllers

import (
	"net/http"
	"users-api/lib/database"
	"users-api/models"

	"github.com/labstack/echo/v4"
)

func CreateUserController(c echo.Context) error {
	var user models.Users
	c.Bind(&user)

	// add input validator?

	if e := database.CreateUser(&user); e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}