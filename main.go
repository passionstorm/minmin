package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
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
	app.Logger().SetLevel("debug")
	err := godotenv.Load()
	if err != nil {
		panic("Not found file .env")
	}
	isDev := os.Getenv("DEV") == "1"
	logic.Load()
	tmpl := iris.Jet("./application/views", ".jet")
	tmpl.Reload(isDev) // remove in production.
	tmpl.AddFunc("base64", func(a view.JetArguments) reflect.Value {
		a.RequireNumOfArguments("base64", 1, 1)
		buffer := bytes.NewBuffer(nil)
		fmt.Fprint(buffer, a.Get(0))
		return reflect.ValueOf(base64.URLEncoding.EncodeToString(buffer.Bytes()))
	})

	app.RegisterView(tmpl)

	//app.HandleDir("/admin", "public/dist/admin", iris.DirOptions{
	//	//Asset:      Asset,
	//	//AssetInfo:  AssetInfo,
	//	//AssetNames: AssetNames,
	//	//ShowList:   true,
	//	IndexName: "/index.html",
	//	// When files should served under compression.
	//	Gzip: true,
	//	// List the files inside the current requested directory if `IndexName` not found.
	//	ShowList: true,
	//})
	//app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	mvcApp := mvc.Configure(app.Party("/api", crs).AllowMethods(iris.MethodOptions), apiRoute)
	mvcApp.HandleError(func(ctx iris.Context, err error) {
		ctx.HTML(fmt.Sprintf("<b>%s</b>", err.Error()))
	})

	staticDomain := app.Party("static.")
	{
		staticDomain.HandleDir("/", "public/dist/admin", iris.DirOptions{
			IndexName: "/index.html",
			Gzip:      true,
			ShowList:  true,
		})
	}

	var port string
	if isDev {
		port = ":9000"
	} else {
		port = ":443"
	}
	err = app.Run(iris.Addr(port))
	die(err)
}

func apiRoute(app *mvc.Application) {
	app.Register(logic.NewUploadLogic())
	app.Register(logic.NewArticleLogic())
	app.Party("/post").Handle(new(R.PostController))
	app.Handle(new(R.UploadController))
}
