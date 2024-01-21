package routes

import (
	"fmt"
	"net/http"

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
	var reqTodo types.CreateTodoRequest
	err := c.Bind(&reqTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("CreateTodo: %+v\n", reqTodo)
	}

	if len(reqTodo.Name) == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "Todo name is required"})
	}

	if reqTodo.UserID == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "Todo userId is required"})
	}

	todo := models.Todo{
		Name:   reqTodo.Name,
		UserID: reqTodo.UserID,
		Done:   reqTodo.Done,
	}

	result := db.DB().Omit("ID").Create(&todo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusCreated, types.ResponseMessage{Message: "Todo created successfully", Data: struct{ ID int64 }{ID: todo.ID}})
}

func DeleteTodo() {}

func CompleteTodo() {}
