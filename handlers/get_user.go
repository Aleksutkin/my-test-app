package handlers

import (
	"my-test-app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type getOneUserServicer interface {
	GetOneUser(int) (models.User, error)
}

func GetUser(gou getOneUserServicer) func(echo.Context) error {
	return func(context echo.Context) error {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			return context.JSON(http.StatusInternalServerError, err.Error())
		}
		usr, err := gou.GetOneUser(id)
		if err != nil {
			return context.JSON(http.StatusNotFound, err.Error())
		}

		return context.JSON(http.StatusOK, usr)
	}
}
