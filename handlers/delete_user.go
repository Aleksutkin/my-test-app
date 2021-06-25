package handlers

import (
	"log"
	"my-test-app/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type deleteUserServicer interface {
	DeleteUser(int) (models.User, error)
}

func DeleteUser(du deleteUserServicer) func(ctx echo.Context) error {
	return func(context echo.Context) error {
		id, err := strconv.Atoi(context.Param("id"))
		if err != nil {
			log.Fatal(err)
		}
		_, err = du.DeleteUser(id)
		if err != nil {
			return context.JSON(http.StatusNotFound, err.Error())
		}

		return context.JSON(http.StatusNoContent, context.NoContent(204))
	}
}
