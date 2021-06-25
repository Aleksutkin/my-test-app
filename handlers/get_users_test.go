package handlers

import (
	"my-test-app/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type handler struct {
	rep *models.Users
	//rep map[string]*models.Users
}

func (h *handler) getUsers(c echo.Context) error {
	users := h.rep
	if users == nil {
		return echo.NewHTTPError(http.StatusNotFound, "user not found")
	}
	return c.JSON(http.StatusOK, users)
}

var (
	rep = &models.Users{
		{Id: 1, Name: "Anton", Email: ""},
		{Id: 2, Name: "Vladimir", Email: "a@a.ra"},
	}

	//rep = map[string]*models.Users{
	//	"1": &models.Users{{Name: "Anton", Email: ""},{Name: "Vladimir", Email: "a@a.ra"}},
	//}

	usersJSON = `{"name":"Anton","email":""},{"name":"Vladimir","email":"a@a.ra"}`
)

func TestGetUsers(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	//c.SetParamNames("id")
	//c.SetParamValues("1")

	h := &handler{rep}

	// Assertions
	if assert.NoError(t, h.getUsers(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, userJSON, rec.Body.String())
	}
}
