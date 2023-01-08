package controllers

import (
	"net/http"
	"task-orm-mvc/models"

	"github.com/labstack/echo/v4"
)

func CreateTiket(c echo.Context) error {
	tiket := models.Tiket{}
	c.Bind(&tiket)

	// send data to database using GORM
	....

	return c.JSON(http.StatusCreated, map[string]any{
		"messages": "success create new tiket",
		"tiket":    tiket,
	})
}

func GetTikets(c echo.Context) error {
	// GORM fecth all data
	....

	return c.JSON(http.StatusFound, map[string]any{
		"messages": "",
		"data":     "data",
	})
}

func GetTiket(c echo.Context) error {
	// GORM fecth a data
	....

	return c.JSON(http.StatusFound, map[string]any{
		"messages": "",
		"data":     "data",
	})
}

func UpdateTiket(c echo.Context) error {
	// GORM update data by id
	....

	return c.JSON(http.StatusCreated, map[string]any{
		"messages": "success create new tiket",
		"tiket":    "tiket",
	})
}

func DeleteTiket(c echo.Context) error {
	// GORM delete data by id
	....
	
	return c.JSON(http.StatusOK, map[string]any{
		"messages": "success create new tiket",
		"tiket":    "tiket",
	})
}
