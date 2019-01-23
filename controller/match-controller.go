package controller

import (
	"net/http"
	"strconv"

	"github.com/Pororift/backend/services"
	"github.com/labstack/echo"
)

// GetMatchByID pororift api for getting a matchID
func GetMatchByID(c echo.Context) error {

	matchID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "given match id is not a number, match id must be a number")
	}

	match, err := services.GetMatch(matchID)

	if err != nil {
		if err.Error() == "403 Forbidden" {
			return c.String(http.StatusForbidden, "api key has expired")
		}
		return c.String(http.StatusNotFound, "match id not found")
	}

	return c.JSON(http.StatusOK, *match)
}

// GetMatchLists pororift api for getting a list of match history
func GetMatchLists(c echo.Context) error {
	accountID := c.Param("id")

	matchList, err := services.GetMatchHistory(accountID)

	if err != nil {
		if err.Error() == "403 Forbidden" {
			return c.String(http.StatusForbidden, "api key has expired")
		}
		return c.String(http.StatusNotFound, "match id not found")
	}

	return c.JSON(http.StatusOK, *matchList)
}
