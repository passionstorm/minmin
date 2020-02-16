package m

import (
	"github.com/kataras/iris/v12"
	"log"
)

func CoverErr(ctx iris.Context) {
	if r := recover(); r != nil {
		ctx.StatusCode(400)
		ctx.JSON(iris.Map{
			"success": 0,
			"error": iris.Map{
				"message": r,
			},
		})
	}
}

func IsEmptyMgo(err error) bool {
	if err == nil {
		return false
	}

	if err.Error() == "mongo: no documents in result" {
		return true
	}

	return false
}
func Err(err error) {
	if err != nil {
		log.Println(err)
	}
}

func HandlErr(err error) {
	if err != nil {
//		tracerr.Print(err)
		panic(err.Error())
	}
}
