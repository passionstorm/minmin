package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"minmin/app/controllers"
)

func NewWebRoute(b *iris.Application) {
	//b.Get("/", func(c iris.Context) {
	//	c.JSON("hello")
	//})
	root := mvc.New(b.Party("/"))
	root.Handle(new(controllers.AdminController))
}
