package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func Logger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		fmt.Printf("Logger: %s\n", ctx.Request().URL.Path)
		return next(ctx)
	}
}
