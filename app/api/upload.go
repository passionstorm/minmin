package api

import (
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"strings"
	"time"
)

func UploadAction(ctx iris.Context) {
	defer coverErr(ctx)
	file, info, err := ctx.FormFile("file")
	handlErr(err)
	defer file.Close()
	now := time.Now()
	fname := strings.Join([]string{
		now.Format("20060102150405"),
		info.Filename,
	}, "_")
	out, err := os.OpenFile("./public/uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
	handlErr(err)
	defer out.Close()
	_, err = io.Copy(out, file)
	handlErr(err)
	ctx.JSON(iris.Map{
		"filename": fname,
	})
	return
}

func coverErr(ctx iris.Context) {
	if r := recover(); r != nil {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"error": iris.Map{
				"message": r,
			},
		})
	}
}
func handlErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
