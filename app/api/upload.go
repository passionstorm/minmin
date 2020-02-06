package api

import (
	"github.com/kataras/iris/v12"
	"io"
	"minmin/app/constant"
	"minmin/app/utils/app"
	"os"
	"strings"
	"time"
)

func UploadAction(ctx iris.Context) {
	defer app.CoverErr(ctx)
	file, info, err := ctx.FormFile("file")
	app.HandlErr(err)
	defer file.Close()
	now := time.Now()
	fname := strings.Join([]string{
		now.Format(constant.YMDHms),
		info.Filename,
	}, "_")
	out, err := os.OpenFile("./public/uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
	app.HandlErr(err)
	defer out.Close()
	_, err = io.Copy(out, file)
	app.HandlErr(err)
	ctx.JSON(iris.Map{
		"filename": fname,
	})
	return
}

