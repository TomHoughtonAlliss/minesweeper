package server

import (
	"minesweeper/controllers"

	"github.com/labstack/echo/v4"
)

func Start() {

	e := echo.New()

	e.GET("/scores", controllers.GetScores)

	e.PUT("/score", controllers.CreateScore)

	e.Logger.Fatal(e.Start(":1323"))
}
