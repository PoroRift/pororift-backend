package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	isDev := false
	if len(os.Args) == 3 {
		isDev = os.Args[1] == "dev"
		os.Setenv("RIOT_KEY", os.Args[2])
	} else {
		fmt.Println("no argument set .... defaulting to prod env")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	if isDev {
		e.GET("api/test", test)
	}

}

/*
	test:
	@param context
	- check to see if connecting to the api works
	- check riot api connection
*/
func test(c echo.Context) error {
	baseURL := os.Getenv("RIOT_API")
	apiKey := os.Getenv("RIOT_KEY")
	res, err := http.Get(fmt.Sprintf("%s/lol/status/v3/shard-data??api_key=%s", baseURL, apiKey))
	if err != nil || res.StatusCode != 200 {
		return c.String(503, "couldn't connect to riot's api at the moment")
	}
	return c.String(200, "connected successfully")
}
