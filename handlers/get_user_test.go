package handlers

import (
	"my-test-app/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type handlerGetUser struct {
	repo map[string]*models.User
}

func (h *handlerGetUser) getUser(c echo.Context) error {
	id := c.Param("id")
	user := h.repo[id]
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, user)
}

var (
	repGetUser = map[string]*models.User{
		"Anton": &models.User{Id: 1, Name: "Anton", Email: "a@aaaaaa.ra"},
	}

	userJSON = `{"name":"Anton","email":"a@aaaaaa.ra"}`
)

func TestGetUser(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:name")
	c.SetParamNames("name")
	c.SetParamValues("Anton")

	h := &handlerGetUser{repGetUser}

	// Assertions
	if assert.NoError(t, h.getUser(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
