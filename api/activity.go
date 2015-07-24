package api

import (
	"net/http"

	"github.com/labstack/echo"
)

//POST /activity
func (a *AppContext) NewActivity(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}

//GET /activities
func (a *AppContext) GetActivities(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}

//GET /activities/:id
func (a *AppContext) FindOneActivity(c *echo.Context) error {

	SetDefaultHeaders(c)

	quote, err := a.Storage.FindOneActivity(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, Error{"Activity could not be found.", err})
	}

	return c.JSON(http.StatusOK, FormatResponse("Fetched", quote))
}

//DELETE /activities/:id
func (a *AppContext) DeleteActivity(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}
