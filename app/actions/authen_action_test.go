package actions

import (
	"fmt"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/mvc"
	"minmin/app/bootstrap"
	"minmin/app/utils/m"
	"testing"
)

func TestAuthenAction_PostRegister(t *testing.T) {
	app := bootstrap.NewApp()
	root := mvc.New(app.Party("/api"))
	root.Handle(new(AuthenAction))

	e := httptest.New(t, app.Application)

	var jsonStr = m.Map{
		"firstname": "minh",
		"lastname": "minh",
		"username": "minhc",
		"password": "qwer1234",
	}
	req := e.POST("/api/register").WithJSON(jsonStr)
	fmt.Println(req.Expect().Body())
}

func TestAuthenAction_PostLogin(t *testing.T) {
	app := bootstrap.NewApp()
	root := mvc.New(app.Party("/api"))
	root.Handle(new(AuthenAction))

	e := httptest.New(t, app.Application)

	var jsonStr = m.Map{
		"username": "minhvh",
		"password": "qwer1234",
	}
	req := e.POST("/api/login").WithJSON(jsonStr)
	fmt.Println(req.Expect().Body())

}
