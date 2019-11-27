package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/_examples/tutorial/mongodb/httputil"
	"minmin/app/models"
)

type PostController struct {
	baseController
}

func (c *PostController) Get() string {

	return "hello"
}
func (c *PostController) GetCreate(ctx iris.Context, logic models.ArticleLogic) {
	_, _ = ctx.JSON(map[string]string{"hello": logic.GetAll()})
}

func (c *PostController) PostCreate(ctx iris.Context, logic models.ArticleLogic) {
	var formData *models.PostForm
	m := new(models.Article)
	err := ctx.ReadJSON(&formData)
	if err != nil {
		_ = httputil.FailJSON(ctx, iris.StatusBadRequest, err, "Param is invalid")
		return
	}
	err = logic.Insert(ctx, formData)
	if err != nil {
		httputil.InternalServerErrorJSON(ctx, err, "Server was unable to create")
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	_, _ = ctx.JSON(m)
}
