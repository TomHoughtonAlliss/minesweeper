package main

import (
	"minesweeper/database"
	"minesweeper/server"
)

func main() {

	database.SetUpDB()

	server.Start()

	defer func () {
		database.CloseConnection()
	}()
}