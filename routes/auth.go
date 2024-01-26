package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"example.com/golang-crud-gorm/db"
	"example.com/golang-crud-gorm/models"
	"example.com/golang-crud-gorm/types"
	"example.com/golang-crud-gorm/utils"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func Register(c echo.Context) error {
	var reqUser types.RegisterRequest
	err := c.Bind(&reqUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, types.ResponseMessage{Message: "Internal Server Error"})
	} else {
		fmt.Printf("CreateUser: %+v\n", reqUser)
	}

	if len(reqUser.UserName) == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "userName is a required field"})
	}

	if len(reqUser.Password) == 0 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "password is a required field"})
	} else if len(reqUser.Password) < 8 {
		return c.JSON(http.StatusBadRequest, types.ResponseMessage{Message: "password must be at least 8 character long"})
	}

	FirstName := reqUser.FirstName
	LastName := reqUser.LastName
	UserName := reqUser.UserName
	Password := reqUser.Password

	Password, _ = utils.HashPassword(Password)

	user := models.User{
		FirstName: sql.NullString{String: FirstName, Valid: true},
		LastName:  sql.NullString{String: LastName, Valid: true},
		UserName:  UserName,
		Password:  Password,
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
