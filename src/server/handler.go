package main

import (
	"net/http"
	"strconv"

	"github.com/PoroRift/pororift-backend/data"
	"github.com/labstack/echo"
)

func test(c echo.Context) error {
	return c.String(http.StatusOK, "This is backend server")
}

func getMatch(c echo.Context) error {
	// name := c.Param("name")
	matchID, _ := strconv.Atoi(c.Param("id"))
	match, _ := data.GetMatch(matchID)
	return c.JSON(http.StatusOK, *match)
}

// func getSummonerStat(c echo.Context) (map[String]string{}) {
// 	// summoner, err := lol.GetSummonerStat(summonerName)
// 	// fmt.Println(summoner)
// 	return nil
// }

// Load static contents
// func getStatic(c echo.Context) error {
// 	return c.String(http.StatusOK, "Getting Static Content")
// }

// Get Champion Rotation
func getChampRot(c echo.Context) error {
	return c.String(http.StatusOK, "Getting Champion Rotation")
}

// Get Summoner information by name
func getSummonerStat(c echo.Context) error {
	// name := c.Param("name")

	// stat, err := lol.SummonerStat(name)
	// fmt.Println(stat)
	// if err == nil {

	// 	return c.JSONPretty(http.StatusOK, stat, " ")
	// }

	// return c.String(http.StatusServiceUnavailable, name)
	// byteSummoner := data.GetPlayer(name)
	return c.String(http.StatusOK, "Getting SummonerStat")
}

func getChampList(c echo.Context) error {
	return c.String(http.StatusOK, "")
}
