package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo"
	st "github.com/wvdeutekom/webhookproject/storage"
)

//POST /activity
func (a *AppContext) NewActivity(c *echo.Context) error {

	return c.JSON(http.StatusOK, FormatResponse("In development", nil))
}

//GET /activities
func (a *AppContext) GetActivities(c *echo.Context) error {

	var activities []st.Activity
	var err error

	var query = c.Request().URL.Query().Get("q")

	//Check for token header

	SetDefaultHeaders(c)

	//Get quote from database
	if query != "" {
		//Seperate search terms and put them into a string array
		activities, err = a.Storage.SearchActivities(strings.Split(query, ","))
	} else {
		activities, err = a.Storage.FindAllActivities()
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, Error{"Activities could not be found.", err})
	}

	return c.JSON(http.StatusOK, FormatResponse("Fetched", activities))
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

	SetDefaultHeaders(c)

	activity, err := a.Storage.DeleteActivity(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, FormatResponse("Deleted", activity))
}
