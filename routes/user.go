package routes

import (
	"fmt"
	"net/http"

	"example.com/golang-crud-gorm/db"
	"example.com/golang-crud-gorm/models"
	"example.com/golang-crud-gorm/types"
	"github.com/labstack/echo/v4"
)

func UpdateUser(c echo.Context) error {
	var user types.UpdateUserRequest

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("UpdateUser: %+v\n", user)
	}

	FirstName := user.FirstName
	LastName := user.LastName

	result := db.DB().Model(&models.User{ID: user.ID}).Select("FirstName", "LastName").Updates(map[string]interface{}{"first_name": FirstName, "last_name": LastName})
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func DeleteUser(c echo.Context) error {
	var user types.DeleteUserRequest

	err := c.Bind(&user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("UpdateUser: %+v\n", user)
	}

	result := db.DB().Delete(&models.User{}, user.ID)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	}

	return c.JSON(http.StatusNoContent, nil)
}
