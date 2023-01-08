package main

import (
	"task-orm-mvc/config"

	"github.com/labstack/echo/v4"
)

func init() {
	config.InitDB()
	config.InitialMigration()
}

func main() {
	// instance of Echo
	e := echo.New()

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
