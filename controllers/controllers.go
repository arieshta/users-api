package controllers

import (
	"net/http"
	"strconv"
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

func GetUsersController(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

func GetUserById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	var user models.Users

	
}

func UpdateUserByIdController(c echo.Context) error {
	var user models.Users
	c.Bind(&user)

	id, _ := strconv.Atoi(c.Param("id"))

	if e := database.UpdateUserById(id, &user); e != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "Success update user",
		"user":   user,
	})
}

func DeleteUserByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	if err := database.DeleteUserById(id); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"message": "record not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success delete user",
	})
}
