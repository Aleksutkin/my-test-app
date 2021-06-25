package handlers

import (
	"log"
	"my-test-app/forms"
	"my-test-app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type PutUsersServicer interface {
	PutUser(forms.User, int) (models.User, error)
}

func PutUser(puu PutUsersServicer) func(echo.Context) error {
	return func(context echo.Context) error {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			log.Fatal(err)
		}

		u := new(forms.User)
		if err := context.Bind(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := context.Validate(u); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		usr, err := puu.PutUser(*u, id)
		if err != nil {
			return context.JSON(http.StatusNotFound, err.Error())
		}

		return context.JSON(http.StatusOK, usr)
	}
}
