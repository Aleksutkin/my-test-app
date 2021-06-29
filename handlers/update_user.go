package handlers

import (
	"my-test-app/forms"
	"my-test-app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UpdateUsersServicer interface {
	UpdateUser(forms.User, int) (*models.User, error)
}

func UpdateUser(puu UpdateUsersServicer) func(echo.Context) error {
	return func(context echo.Context) error {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			return context.JSON(http.StatusNotFound, err.Error())
		}

		u := new(forms.User)
		if err := context.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := context.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		usr, err := puu.UpdateUser(*u, id)
		if err != nil && usr == nil {
			return context.JSON(http.StatusNotFound, err.Error())
		}
		if err != nil && usr != nil {
			return context.JSON(http.StatusInternalServerError, err.Error())
		}

		return context.JSON(http.StatusOK, usr)
	}
}
