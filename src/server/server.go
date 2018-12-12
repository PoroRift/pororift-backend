package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	// fmt.Println(lol.Test())
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// CORS
	// gives web servers cross-domain access controls, which enable secure cross-domain data transfer
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	e.GET("/", test)

	e.GET("/summoner/:name", getSummonerStat)

	e.GET("/champ_rotation", getChampRot)

	e.GET("/champ", getChampList)

	// TODO: Need to make list of all avai champions art
	// TODO: remove relative path
	e.Static("/static/champion", "../data/img/champion")

	// e.GET("/static/avail_art" )

	e.Logger.Fatal(e.Start(":3001"))
}
