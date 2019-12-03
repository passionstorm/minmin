package controllers

import (
	"github.com/kataras/iris/v12"
	"minmin/app/logic"
	"minmin/app/pkg/httputil"
)

type PostController struct {
	ApiController
}

func (c *PostController) GetEditBy(id int, ctx iris.Context, l logic.ArticleLogic) {
	rs, err := l.FindById(id)
	if err != nil {
		_ = httputil.FailJSON(ctx, iris.StatusNotFound, err, "Param is invalid")
		return
	}
	_, _ = ctx.JSON(&rs)
}

func (c *PostController) Get(ctx iris.Context, l logic.ArticleLogic) {
	ctx.JSON(l.GetAll())
}

func (c *PostController) PostCreate(ctx iris.Context, db logic.ArticleLogic) {
	var formData *logic.PostForm
	err := ctx.ReadJSON(&formData)
	if err != nil {
		_ = httputil.FailJSON(ctx, iris.StatusBadRequest, err, "Param is invalid")
		return
	}
	err = db.Insert(ctx, formData)
	if err != nil {
		httputil.InternalServerErrorJSON(ctx, err, "Server was unable to create")
		return
	}

	c.Success(iris.StatusOK, nil)
}
