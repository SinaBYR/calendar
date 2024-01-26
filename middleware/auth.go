package middleware

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Authenticate() echo.MiddlewareFunc {
	return echojwt.JWT([]byte("secret"))
}
