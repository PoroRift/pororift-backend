package controller

import (
	"net/http"
	"strconv"

	"github.com/Pororift/backend/services"
	"github.com/labstack/echo"
)

func GetMatchById(c echo.Context) error {

	matchID, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return c.String(http.StatusBadRequest, "given match id is not a number, match id must be a number")
	}

	match, err := services.GetMatch(matchID)
	if err != nil {
		return c.String(http.StatusNotFound, "match id not found")
	}

	return c.JSON(http.StatusOK, match)
}

func GetMatchLists(c echo.Context) error {
	return c.String(http.StatusOK, "history list")
}
