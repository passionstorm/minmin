package routes

import (
	"github.com/kataras/iris/v12"
	"minmin/app/bootstrap"
)

func NewWebRoute(b *bootstrap.Rooter) {
	b.Get("/", func(c iris.Context) {
		c.ViewData("title", "Hello")
		c.View("home.jet")
	})
}
