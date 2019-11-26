package controllers

import (
	"github.com/kataras/iris/v12"
	"html"
	"minmin/app/models"
)

type PostController struct {
	baseController
	Model models.Article
}

func (c *PostController) Get() string {

	return "hello"
}
func (c *PostController) GetCreate(ctx iris.Context) {
	c.Model.InitData()
	_, _ = ctx.JSON(map[string]string{"hello": "json"})
}

type PostForm struct {
	Title      string
	Content    string
	Categories []string
}

func (c *PostController) PostCreate(ctx iris.Context) {
	form := PostForm{}
	_ = c.Ctx.ReadJSON(&form)
	//fmt.Println(form)
	form.Content = html.EscapeString(form.Content)

	_, _ = ctx.JSON(form)
}
