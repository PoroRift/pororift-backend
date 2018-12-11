package main

import (
	"fmt"

	"github.com/PoroRift/pororift-backend/lol"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	fmt.Println(lol.Test())
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

	e.GET("/static/:category/:championid", getStatic)

	e.Logger.Fatal(e.Start(":3001"))
}
