package actions

import "github.com/labstack/echo/v4"

type BaseAction interface {
	Routes(c *echo.Echo)
}
