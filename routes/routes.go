package routes

import (
	"task-orm-mvc/controllers"

	"github.com/labstack/echo/v4"
)

func Init() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// Route of User / to handler function
	e.GET("/users", controllers.GetUsersController)
	e.GET("/users/:id", controllers.GetUserController)
	e.POST("/users", controllers.CreateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)
	e.PUT("/users/:id", controllers.UpdateUserController)

	// Route of tiket / to handler function
	e.GET("/tikets", ....)
	e.GET("/tikets/:id", ....)
	e.POST("/tikets", ....)
	e.DELETE("/tikets/:id", ....)
	e.PUT("/tikets/:id", ....)

	// Route of transaksi / to handler function
	....

	return e
}
