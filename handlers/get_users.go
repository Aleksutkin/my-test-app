package handlers

import (
	"my-test-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type getAllUsersServicer interface {
	GetAllUsers() (models.Users, error)
}

func GetUsers(gau getAllUsersServicer) func(echo.Context) error {
	return func(context echo.Context) error {
		usr, err := gau.GetAllUsers()
		if err != nil {
			return context.JSON(http.StatusInternalServerError, err.Error())
		}

		return context.JSON(http.StatusOK, usr)
	}
}
