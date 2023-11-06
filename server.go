package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func play(c echo.Context) {

}

func main() {
	e := echo.New()
	e.GET("/play", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello mate")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
