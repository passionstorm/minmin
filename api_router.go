package main

import (
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

func initApiDomain(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	mvcApp := mvc.Configure(app.Subdomain("api", crs).AllowMethods(iris.MethodOptions), apiRoute)
	mvcApp.HandleError(func(ctx iris.Context, err error) {
		ctx.HTML(fmt.Sprintf("<b>%s</b>", err.Error()))
	})
}

