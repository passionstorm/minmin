package actions

import (
	"fmt"
	"github.com/kataras/iris/v12/httptest"
	"github.com/kataras/iris/v12/mvc"
	"minmin/app/bootstrap"
	"minmin/app/utils/m"
	"testing"
)

var app *bootstrap.Rooter

func init()  {
	app = bootstrap.NewApp()
	root := mvc.New(app.Party("/api"))
	root.Handle(new(AuthenAction))
}

func TestAuthenAction_PostRegister(t *testing.T) {
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
	e := httptest.New(t, app.Application)

	var jsonStr = m.Map{
		"username": "minhvh",
		"password": "qwer1234",
	}
	req := e.POST("/api/login").WithJSON(jsonStr)
	fmt.Println(req.Expect().Body())

}

func TestAuthenAction_GetInfo(t *testing.T) {
	e := httptest.New(t, app.Application)
	req := e.GET("/api/info").WithHeader("Authorization", `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJmaXJzdG5hbWUiOiJtaW5oIiwibGFzdG5hbWUiOiJtaW5oIiwidXNlcm5hbWUiOiJtaW5odmgifQ.hZMkEUp3pjVNPy29Mft0Zj_xupuwIKekYE5junS9Uv4`)
	fmt.Println(req.Expect().Body())
}