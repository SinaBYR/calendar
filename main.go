package main

import (
	"example.com/golang-crud-gorm/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// err := db.Migrate()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// db.DB().AutoMigrate(&models.Todo{})

	e := echo.New()

	e.POST("/users/create", routes.CreateUser)
	e.PUT("/users/:id", routes.UpdateUser)
	e.DELETE("/users/:id", routes.DeleteUser)

	e.POST("/todos/create", routes.CreateTodo)
	e.DELETE("/todos/:id", routes.DeleteTodo)

	e.Logger.Fatal(e.Start(":8000"))
}
