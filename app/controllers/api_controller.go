package controllers

import (
	"github.com/kataras/iris/v12"
)

type ApiController struct {
	Ctx iris.Context
}

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func (c *ApiController) Success(statusCode int, data interface{}) {
	if statusCode == 0 {
		c.Ctx.StatusCode(iris.StatusOK)
	} else {
		c.Ctx.StatusCode(statusCode)
	}

	_, _ = c.Ctx.JSON(&Response{Success: true, Data: data})

}
