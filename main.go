package main

import (
	"github.com/kataras/iris/v12"
	"minmin/app/bootstrap"
	"minmin/app/routes"
)


func main() {
	app := bootstrap.NewApp()
	app.Application.Configure(
		routes.NewApiRoute,
		routes.NewWebRoute,
	)
	app.Listen(
		":9000",
		// Ignores err server closed log when CTRL/CMD+C pressed.
		//iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
