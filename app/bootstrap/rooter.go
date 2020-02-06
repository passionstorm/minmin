package bootstrap

import (
	"github.com/gorilla/securecookie"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/sessions"
	"github.com/kataras/iris/v12/websocket"
	"log"
	"time"
)

type Configurator func(*Rooter)

type Rooter struct {
	*iris.Application
	AppName      string
	AppOwner     string
	AppSpawnDate time.Time

	Sessions *sessions.Sessions
}

func New(appName, appOwner string, cfgs ...Configurator) *Rooter {
	b := &Rooter{
		AppName:      appName,
		AppOwner:     appOwner,
		AppSpawnDate: time.Now(),
		Application:  iris.New(),
	}

	for _, cfg := range cfgs {
		cfg(b)
	}

	return b
}



// SetupSessions initializes the sessions, optionally.
func (b *Rooter) SetupSessions(expires time.Duration, cookieHashKey, cookieBlockKey []byte) {
	b.Sessions = sessions.New(sessions.Config{
		Cookie:   "SECRET_SESS_COOKIE_" + b.AppName,
		Expires:  expires,
		Encoding: securecookie.New(cookieHashKey, cookieBlockKey),
	})
}

// SetupWebsockets prepares the websocket server.
func (b *Rooter) SetupWebsockets(endpoint string, handler websocket.ConnHandler) {
	ws := websocket.New(websocket.DefaultGorillaUpgrader, handler)

	b.Get(endpoint, websocket.Handler(ws))
}

// SetupErrorHandlers prepares the http error handlers
// `(context.StatusCodeNotSuccessful`,  which defaults to < 200 || >= 400 but you can change it).
func (b *Rooter) SetupErrorHandlers() {
	b.OnAnyErrorCode(func(c iris.Context) {
		err := iris.Map{
			"app":     b.AppName,
			"status":  c.GetStatusCode(),
			"message": c.Values().GetString("message"),
		}

		if jsonOutput := c.URLParamExists("json"); jsonOutput {
			c.JSON(err)
			return
		}

		c.ViewData("Err", err)
		c.ViewData("Title", "Error")
		c.View("shared/error.jet")
	})
}

const (
	StaticAssets = "./public/"
	Favicon      = "favicon.ico"
)

// Configure accepts configurations and runs them inside the Bootstraper's context.
func (b *Rooter) Configure(cs ...Configurator) {
	for _, c := range cs {
		c(b)
	}
}

// Bootstrap prepares our application.
//
// Returns itself.
func (b *Rooter) Bootstrap() *Rooter {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	b.SetupViews("./app/views")
	b.SetupSessions(24*time.Hour,
		[]byte("the-big-and-secret-fash-key-here"),
		[]byte("lot-secret-of-characters-big-too"),
	)
	b.SetupErrorHandlers()

	// middleware, after static files
	b.Use(recover2.New())
	b.Use(logger.New())

	return b
}

// Listen starts the http server with the specified "addr".
func (b *Rooter) Listen(addr string, cfgs ...iris.Configurator) {
	b.Run(iris.Addr(addr), cfgs...)
}
