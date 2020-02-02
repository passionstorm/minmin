package api

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
)

func NewApiDomain(app *iris.Application) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "PUT","POST", "DELETE", "HEAD", "OPTIONS"},
		AllowCredentials: true,
	})
	route := app.Subdomain("api", crs)
	{
		route.Get("/", func(ctx iris.Context) {
			ctx.JSON(map[string]interface{}{
				"google": "google",
			})
		})
		route.Post("/upload", UploadAction)
	}
}


