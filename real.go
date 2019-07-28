package realgo

import (
	"github.com/julienschmidt/httprouter"
	"realgo/conf"
	"net/http"
	"fmt"
)



type App struct {
	Config *AppConfig
	Server *WebServer
}
type AppConfig struct {
	AppName  string `toml:"app_name"`
	RunMode  string `toml:"run_mode"`
	Host     string `toml:"host"`
	Port     string `toml:"port"`
}

func New() *App{
	appconfig := &AppConfig{}
	conf.ReadAppConfFile("app.toml", appconfig)
	webServer := &WebServer{
		Router:httprouter.New(),
	}
	engine := &App{
		Config: appconfig,
		Server: webServer,
	}
	return  engine
}

func Test()  *httprouter.Router{
	return httprouter.New()
}

type RouterRegister func(*WebServer)

func (app *App)WebServer() *WebServer {
	return app.Server
}

func (app *App) RegisterRouter(rr RouterRegister)  {
	rr(app.WebServer())
}

func (app *App)StartServer()  {
	fmt.Println(app.Config.Host+":"+app.Config.Port)

	http.ListenAndServe(app.Config.Host+":"+app.Config.Port, app.WebServer().Router)
}


