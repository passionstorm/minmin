package middlewares

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

func MRecover() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			defer func() {
				if r := recover(); r != nil {
					err, ok := r.(error)
					if !ok {
						err = fmt.Errorf("%v", r)
					}
					c.Error(err)
				}
			}()
			return next(c)
		}
	}
}
