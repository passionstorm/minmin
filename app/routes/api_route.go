package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"minmin/app/actions"
)

func NewApiRoute(b *iris.Application) {
	//crs := cors.New(cors.Options{
	//	AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
	//	AllowedHeaders:   []string{"*"},
	//	AllowedMethods:   []string{"GET", "PUT","POST", "DELETE", "HEAD", "OPTIONS"},
	//	AllowCredentials: true,
	//})
	root := mvc.New(b.Party("/api"))
	root.Handle(new(actions.AuthenAction))
	//
	//route := b.Subdomain("api", crs)
	//{
	//	route.Get("/", func(ctx iris.Context) {
	//		ctx.JSON(map[string]interface{}{
	//			"google": "google",
	//		})
	//	})
	//	route.Post("/upload", r.UploadAction)
	//}
}
