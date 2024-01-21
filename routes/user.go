package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/golang-crud-gorm/db"
	"example.com/golang-crud-gorm/models"
	"example.com/golang-crud-gorm/types"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	var reqUser types.CreateUserRequest
	err := c.Bind(&reqUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("CreateUser: %+v\n", reqUser)
	}

	if len(reqUser.UserName) == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "userName is a required field"})
	}

	FirstName := reqUser.FirstName
	LastName := reqUser.LastName
	UserName := reqUser.UserName

	user := models.User{
		FirstName: sql.NullString{String: FirstName, Valid: true},
		LastName:  sql.NullString{String: LastName, Valid: true},
		UserName:  UserName,
	}

	result := db.DB().Omit("ID").Create(&user)

	if result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1062:
				{
					return c.JSON(http.StatusConflict, types.ResponseMessage{Message: fmt.Sprintf("A user with username '%v' already exists", reqUser.UserName)})
				}
			default:
				return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
			}
		}
	}

	return c.JSON(http.StatusCreated, types.ResponseMessage{Message: "User created successfully", Data: struct{ ID int64 }{ID: user.ID}})
}

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
