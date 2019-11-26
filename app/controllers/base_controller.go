package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/sessions"
)

type baseController struct {
	Ctx iris.Context
	Session *sessions.Session
}


