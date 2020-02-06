package app

import "github.com/kataras/iris/v12"

func CoverErr(ctx iris.Context) {
	if r := recover(); r != nil {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"error": iris.Map{
				"message": r,
			},
		})
	}
}
func HandlErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
