package main

import (
	"minesweeper/handler"

	"github.com/labstack/echo/v4"
)

func play(c echo.Context) {

}

func main() {

	handler := handler.Handler{}

	e := echo.New()
	e.GET("/scores", handler.GetRecords)
	e.Logger.Fatal(e.Start(":1323"))
}
