package routes

import (
	"fmt"
	"net/http"
	"time"

	"example.com/golang-crud-gorm/db"
	"example.com/golang-crud-gorm/models"
	"example.com/golang-crud-gorm/types"
	"github.com/labstack/echo/v4"
)

func GetTodo(c echo.Context) {}

// func hello(c echo.Context) error {
// 	return c.String(http.StatusOK, "Hello World")
// }

func GetTodos() {}

func CreateTodo(c echo.Context) error {
	todo := models.Todo{Done: false, CreatedAt: time.Now().UnixMilli()}
	err := c.Bind(&todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("CreateTodo: %+v\n", todo)
	}

	if len(todo.Name) == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "Todo name is required"})
	}

	fmt.Printf("%d\n", todo.UserID)
	if todo.UserID == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "Todo userId is required"})
	}

	result := db.DB().Omit("ID").Create(&todo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusCreated, types.ResponseMessage{Message: "Todo created successfully", Data: struct{ ID int64 }{ID: todo.ID}})
}

func DeleteTodo() {}

func CompleteTodo() {}
