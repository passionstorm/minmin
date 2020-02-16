package m

import (
	"errors"
	"github.com/kataras/iris/v12"
	"github.com/ztrue/tracerr"
	"log"
)

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

func HandlErr(err error, newMsg ...string) {
	if err != nil {
		tracerr.PrintSourceColor(err, 2)
		if len(newMsg) > 0 {
			panic(errors.New(newMsg[0]).Error())
		}
		//
		panic(err.Error())
	}
}
