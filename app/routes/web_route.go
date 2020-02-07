package routes

import (
	"github.com/kataras/iris/v12/mvc"
	"minmin/app/bootstrap"
	"minmin/app/controllers"
)

func NewWebRoute(b *bootstrap.Rooter) {
	//b.Get("/", func(c iris.Context) {
	//	c.JSON("hello")
	//})
	root := mvc.New(b.Party("/"))
	root.Handle(new(controllers.AdminController))
}
