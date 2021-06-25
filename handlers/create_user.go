package handlers

import (
	"my-test-app/forms"
	"my-test-app/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateUsersServicer interface {
	CreateUser(forms.User) (models.User, error)
}

func CreateUser(pou CreateUsersServicer) func(echo.Context) error {
	return func(context echo.Context) error {
		u := new(forms.User)
		if err := context.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := context.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		usr, err := pou.CreateUser(*u)
		if err != nil {
			return context.JSON(http.StatusInternalServerError, err.Error())
		}

		return context.JSON(http.StatusOK, usr)
	}
}
