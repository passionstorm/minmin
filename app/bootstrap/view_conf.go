package bootstrap

import (
	"github.com/kataras/iris/v12"
	"reflect"
	"strings"
)

func (b *Rooter) SetupViews(viewsDir string) {
	tmpl := iris.Jet(viewsDir, ".jet") // <--
	tmpl.Reload(true)
	val := reflect.ValueOf(ViewBuiler{})
	fns := val.Type()
	for i := 0; i < fns.NumMethod(); i++ {
		tmpl.AddFunc(strings.ToLower(fns.Method(i).Name), val.Method(i).Interface())
	}

	b.RegisterView(tmpl)
}
