package main

import (
	"fmt"
	"net/http"

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

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This is backend server")
	})

	e.Logger.Fatal(e.Start(":3001"))
}
