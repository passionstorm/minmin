package main

import (
	"github.com/kataras/iris/v12"
	"minmin/app/bootstrap"
	"minmin/app/routes"
)

func newApp() *bootstrap.Rooter {
	t := bootstrap.New("Mintoot", "vohoangminhdn93@gmail.com")
	t.Bootstrap()
	t.Configure(
		routes.NewApiRoute,
		routes.NewWebRoute,
	)

	return t
}

func main() {
	app := newApp()
	app.Listen(
		"mintoot.local:80",
		// Ignores err server closed log when CTRL/CMD+C pressed.
		//iris.WithoutServerError(iris.ErrServerClosed),
		// Enables faster json serialization and more.
		iris.WithOptimizations,
	)
}
