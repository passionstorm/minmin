package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type AdminController struct {
	baseController
}

func (c *AdminController) GetIndex() mvc.Result {
	return mvc.View{
		Name: "login.jet",
		Data: map[string]interface{}{
			"title": "Hello Page",
			"":      "Welcome to my awesome website",
		},
	}
}

//admin/login
func (c *AdminController) GetLogin() mvc.Result {
	return mvc.View{
		Name: "login.jet",
		Data: map[string]interface{}{
			"title": "Hello Page",
			"":      "Welcome to my awesome website",
		},
	}
}

func (c *AdminController) PostLogin(ctx iris.Context) mvc.Result {
	username := ctx.PostValueTrim("username")
	password := ctx.PostValueTrim("password")
	title := "Login"
	if username == "admin" && password == "admin" {
		title = "Login success"
	}

	return mvc.View{
		Name: "login.jet",
		Data: map[string]interface{}{
			"title": title,
			"":      "Welcome to my awesome website",
		},
	}
}

