package main

import (
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
	app.Listen(":9000")
}

