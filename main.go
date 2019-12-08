package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/view"
	R "minmin/app/controllers"
	"minmin/app/logic"
	"os"
	"reflect"
)

func die(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	app := iris.New()
	app.Use(recover.New())
	//app.Use(logger.New())
	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		ctx.HTML("Message: <b>" + ctx.Values().GetString("message") + "</b>")
	})
	app.Logger().SetLevel("debug")
	err := godotenv.Load()
	if err != nil {
		panic("Not found file .env")
	}
	isDev := os.Getenv("DEV") == "1"
	logic.Load()
	tmpl := iris.Jet("./app/views", ".jet")
	tmpl.Reload(isDev) // remove in production.
	tmpl.AddFunc("base64", func(a view.JetArguments) reflect.Value {
		a.RequireNumOfArguments("base64", 1, 1)
		buffer := bytes.NewBuffer(nil)
		fmt.Fprint(buffer, a.Get(0))
		return reflect.ValueOf(base64.URLEncoding.EncodeToString(buffer.Bytes()))
	})

	app.RegisterView(tmpl)

	initApiDomain(app)
	initAdminDomain(app)
	var port string
	if isDev {
		port = ":9000"
	} else {
		port = ":443"
	}
	err = app.Run(iris.Addr(port))
	die(err)
}

func registerLogic(app *mvc.Application) {
	app.Register(logic.NewUploadLogic())
	app.Register(logic.NewArticleLogic())
}

func apiRoute(app *mvc.Application) {
	registerLogic(app)

	app.Party("/post").Handle(new(R.PostController))
	app.Handle(new(R.UploadController))
}
