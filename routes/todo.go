package routes

import (
	"errors"
	"fmt"
	"net/http"

	"example.com/golang-crud-gorm/db"
	"example.com/golang-crud-gorm/models"
	"example.com/golang-crud-gorm/types"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetTodo(c echo.Context) error {
	var reqTodo types.GetTodoRequest

	err := c.Bind(&reqTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Errorr"})
	} else {
		fmt.Printf("GetTodo: %+v\n", reqTodo)
	}

	if reqTodo.ID == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "Todo id is required"})
	}

	var todo = types.Todo{}

	result := db.DB().Where(&types.Todo{ID: reqTodo.ID}).First(&todo)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, types.ResponseMessage{Message: fmt.Sprintf("resource not found")})
	}

	return c.JSON(http.StatusOK, types.ResponseMessage{
		Data: todo,
	})
}

func GetTodos(c echo.Context) error {
	var reqTodo types.GetTodosRequest
	err := c.Bind(&reqTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("GetTodos: %+v\n", reqTodo)
	}

	var todos []types.Todo

	result := db.DB().Find(&todos)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusOK, types.ResponseMessage{
		Data: todos,
	})
}

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

	todo := types.Todo{
		Name: reqTodo.Name,
		Done: reqTodo.Done,
	}

	result := db.DB().Omit("ID").Create(&todo)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusCreated, types.ResponseMessage{Message: "Todo created successfully", Data: todo})
}

func DeleteTodo(c echo.Context) error {
	var reqTodo types.DeleteTodoRequest

	err := c.Bind(&reqTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("DeleteTodo: %v\n", reqTodo)
	}

	if reqTodo.ID == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "Todo id is required"})
	}

	result := db.DB().Delete(&types.Todo{}, reqTodo.ID)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func CompleteTodo(c echo.Context) error {
	var reqTodo types.CompleteTodoRequest

	err := c.Bind(&reqTodo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("CompleteTodo: %v\n", reqTodo)
	}

	if reqTodo.ID == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "Todo id is required"})
	}

	result := db.DB().Model(&models.Todo{ID: reqTodo.ID}).Select("Done").Updates(map[string]interface{}{"done": true})

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusNoContent, nil)
}
