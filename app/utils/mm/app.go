package mm

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/ztrue/tracerr"
	"log"
	"net/http"
)

func CoverErr(c echo.Context) {
	if r := recover(); r != nil {
		c.JSON(http.StatusBadRequest, Map{
			"error": Map{
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

func FindStr(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
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
