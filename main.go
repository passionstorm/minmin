package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/view"
	R "minmin/app/controllers"
	"minmin/app/logic"
	"reflect"
)

func main() {
	app := iris.New()
	app.Use(recover.New())
	app.Logger().SetLevel("debug")

	tmpl := iris.Jet("./application/views", ".jet")
	tmpl.Reload(true) // remove in production.
	tmpl.AddFunc("base64", func(a view.JetArguments) reflect.Value {
		a.RequireNumOfArguments("base64", 1, 1)
		buffer := bytes.NewBuffer(nil)
		fmt.Fprint(buffer, a.Get(0))
		return reflect.ValueOf(base64.URLEncoding.EncodeToString(buffer.Bytes()))
	})
	app.RegisterView(tmpl)

	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin/login") })

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	mvcApp := mvc.Configure(app.Party("/api", crs).AllowMethods(iris.MethodOptions), apiRoute)

	mvcApp.HandleError(func(ctx iris.Context, err error) {
		ctx.HTML(fmt.Sprintf("<b>%s</b>", err.Error()))
	})

	err := app.Run(iris.Addr(":9000"))
	if err != nil {
		panic(err)
	}
}

func notFound() {

}

func apiRoute(app *mvc.Application) {
	app.Register(logic.NewUploadLogic())
	app.Register(logic.NewArticleLogic()).Party("/post").Handle(new(R.PostController))
	app.Handle(new(R.UploadController))
}

// localstion/admin/...
