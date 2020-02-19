package main

import (
	"minmin/app/actions"
	"minmin/app/bootstrap"
)

func init() {
	app := bootstrap.NewApp()
	app.ImportRoutes(actions.AuthenAction{})
	app.Debug = true
	app.Logger.Print(app.Routes())
	app.Listen(":9999")

}

func main() {


}
