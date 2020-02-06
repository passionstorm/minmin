package routes

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	r "minmin/app/api"
	"minmin/app/bootstrap"
)

func NewApiRoute(b *bootstrap.Rooter) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT","POST", "DELETE", "HEAD", "OPTIONS"},
		AllowCredentials: true,
	})
	route := b.Subdomain("api", crs)
	{
		route.Get("/", func(ctx iris.Context) {
			ctx.JSON(map[string]interface{}{
				"google": "google",
			})
		})
		route.Post("/upload", r.UploadAction)
	}
}
