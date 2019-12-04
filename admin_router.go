package main

import (
	"github.com/kataras/iris/v12"
	"os"
)

func initAdminDomain(app *iris.Application) {
	adDomain := app.Subdomain("admin")
	{
		adDomain.HandleDir("/", "public/dist/admin", iris.DirOptions{
			IndexName: "/index.html",
			Gzip:      true,
			ShowList:  true,
		})
		adminHost := "http://" + os.Getenv("ADMIN_HOST")
		adDomain.Get("/dashboard", func(ctx iris.Context) {
			ctx.View("admin/admin.jet", map[string]interface{}{
				"AdminHost": adminHost,
			})
		})
		adDomain.Get("/", func(ctx iris.Context) {
			ctx.View("admin/admin.jet", map[string]interface{}{
				"AdminHost": adminHost,
			})
		})
		adDomain.Get("/member", func(ctx iris.Context) {
			ctx.View("admin/admin.jet", map[string]interface{}{
				"AdminHost": adminHost,
			})
		})
		adDomain.Get("/member/permission", func(ctx iris.Context) {
			ctx.View("admin/admin.jet", map[string]interface{}{
				"AdminHost": adminHost,
			})
		})
		adDomain.Get("/post", func(ctx iris.Context) {
			ctx.View("admin/admin.jet", map[string]interface{}{
				"AdminHost": adminHost,
			})
		})
		adDomain.Get("/post/edit", func(ctx iris.Context) {
			ctx.View("admin/admin.jet", map[string]interface{}{
				"AdminHost": adminHost,
			})
		})
	}
}
