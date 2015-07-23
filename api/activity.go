package api

import (
	"net/http"

	"github.com/labstack/echo"
)

//POST /activity
func (a *AppContext) NewActivity(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}

//GET /activity
func (a *AppContext) GetActivities(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}

//DELETE /activity/:id
func (a *AppContext) DeleteActivity(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}
