package main

import (
	"example.com/golang-crud-gorm/middleware"
	"example.com/golang-crud-gorm/routes"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
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

	todoRouter := e.Group("/todos")
	todoRouter.Use(middleware.Logger)

	todoRouter.GET("", routes.GetTodos, echoMiddleware.AddTrailingSlash()) // it won't work if you add slash
	todoRouter.GET("/:id", routes.GetTodo)
	todoRouter.POST("/create", routes.CreateTodo)
	todoRouter.PUT("/:id", routes.CompleteTodo)
	todoRouter.DELETE("/:id", routes.DeleteTodo)

	e.Logger.Fatal(e.Start(":8000"))
}
