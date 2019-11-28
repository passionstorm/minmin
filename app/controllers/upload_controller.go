package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/_examples/tutorial/mongodb/httputil"
	"minmin/app/logic"
)

type UploadController struct {
	ApiController
}

func (c *UploadController) PostUpload(ctx iris.Context, l logic.UploadLogic) {
	err := l.UpFile(ctx)
	if err != nil {
		httputil.InternalServerErrorJSON(ctx, err, "Server can't upload")
		return
	}
}

func (c *UploadController) GetUpload(ctx iris.Context, l logic.UploadLogic) {
	items := l.GetUploads()
	c.Success(iris.StatusOK, items)
}
