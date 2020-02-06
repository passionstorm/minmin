package bootstrap

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/kataras/iris/v12/view"
	"os"
	"reflect"
)

type ViewBuiler struct {
}

func (ViewBuiler) Asset(a view.JetArguments) reflect.Value {
	path := a.Get(0).String()
	fmt.Println(os.Getenv("APP_URL"))
	return reflect.ValueOf(path)
}

func (ViewBuiler) Style(a view.JetArguments) reflect.Value {
	path := a.Get(0).String()
	s := fmt.Sprintf(`<link href="%v" rel="stylesheet"> `, path)
	return reflect.ValueOf(s)
}

func (ViewBuiler) Script(a view.JetArguments) reflect.Value {
	path := a.Get(0).String()
	s := fmt.Sprintf(`<script src="%v" ></script>`, path)
	return reflect.ValueOf(s)
}

func (ViewBuiler) Base64(a view.JetArguments) reflect.Value {
	a.RequireNumOfArguments("base64", 1, 1)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprint(buffer, a.Get(0))
	return reflect.ValueOf(base64.URLEncoding.EncodeToString(buffer.Bytes()))
}
