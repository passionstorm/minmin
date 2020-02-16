package api

import (
	"github.com/kataras/iris/v12"
	"io"
	"minmin/app/constant"
	"minmin/app/utils/m"
	"os"
	"strings"
	"time"
)

func UploadAction(ctx iris.Context) {
	defer m.CoverErr(ctx)
	file, info, err := ctx.FormFile("file")
	m.HandlErr(err)
	defer file.Close()
	now := time.Now()
	fname := strings.Join([]string{
		now.Format(constant.YMDHms),
		info.Filename,
	}, "_")
	out, err := os.OpenFile("./public/uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
	m.HandlErr(err)
	defer out.Close()
	_, err = io.Copy(out, file)
	m.HandlErr(err)
	ctx.JSON(m.Map{
		"filename": fname,
	})
	return
}

