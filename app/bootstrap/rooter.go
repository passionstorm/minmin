package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/ztrue/tracerr"
	"minmin/app/actions"
	"time"
)

type Configurator func(*Rooter)

type Rooter struct {
	*echo.Echo
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time
}

func New(appName, appOwner string) *Rooter {
	e := &Rooter{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Echo:         echo.New(),
	}

	return e
}

func (e *Rooter) ImportRoutes(cs ...actions.BaseAction) {
	for _, c := range cs {
		c.Routes(e.Echo)
	}
}

func (e *Rooter) errorHandler(err error, c echo.Context) {
	tracerr.Print(err)
	e.DefaultHTTPErrorHandler(err, c)
}

const (
	StaticAssets = "./spa/dist"
	AppViews     = "./app/views"
)

// Bootstrap prepares our application.
//
// Returns itself.
func (e *Rooter) Bootstrap() *Rooter {
	//err := godotenv.Load()
	//m.HandlErr(tracerr.Wrap(err))
	//e.SetupViews(AppViews)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339}[${method}][${status}] ${uri} (${latency_human})\n",
	}))
	e.Use(middleware.Recover())
	//e.Use(middleware.Static(StaticAssets))
	e.HTTPErrorHandler = e.errorHandler

	return e
}

func NewApp() *Rooter {
	t := New("Mintoot", "vohoangminhdn93@gmail.com")
	t.Bootstrap()
	return t
}

// Listen starts the http server with the specified "addr".
func (e *Rooter) Listen(addr string) {
	e.Logger.Fatal(e.Start(addr))
}
