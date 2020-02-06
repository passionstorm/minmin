package bootstrap

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/view"
	"os"
	"reflect"
	"strings"
)

func fnAsset(a view.JetArguments) reflect.Value {
	path := a.Get(0).String()
	fmt.Println(os.Getenv("APP_URL"))
	return reflect.ValueOf(path)
}

func fnStyle(a view.JetArguments) reflect.Value {
	path := a.Get(0).String()
	s := fmt.Sprintf(`<link href="%v" rel="stylesheet"> `, path)
	return reflect.ValueOf(s)
}

func fnScript(a view.JetArguments) reflect.Value {
	path := a.Get(0).String()
	s := fmt.Sprintf(`<script src="%v" ></script>`, path)
	return reflect.ValueOf(s)
}

func fnBase64(a view.JetArguments) reflect.Value {
	a.RequireNumOfArguments("base64", 1, 1)
	buffer := bytes.NewBuffer(nil)
	fmt.Fprint(buffer, a.Get(0))
	return reflect.ValueOf(base64.URLEncoding.EncodeToString(buffer.Bytes()))
}

func (b *Rooter) SetupViews(viewsDir string) {
	tmpl := iris.Jet(viewsDir, ".jet") // <--
	tmpl.Reload(true)

	clz := ViewBuiler{}
	fns := reflect.TypeOf(ViewBuiler{})
	for i := 0; i < fns.NumMethod(); i++ {
		method := fns.Method(i)
		tmpl.AddFunc(strings.ToLower(method.Name), reflect.ValueOf(&clz).Method(i).)
	}

	b.RegisterView(tmpl)
}
