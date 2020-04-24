package realgo

import (
	"github.com/greatming/realgo/conf"
	"github.com/greatming/realgo/lib/graceful"
	"github.com/julienschmidt/httprouter"
	"log"
	"os"
)

type App struct {
	Config *AppConfig
	Server *WebServer
}
type AppConfig struct {
	AppName string `toml:"app_name"`
	RunMode string `toml:"run_mode"`
	Host    string `toml:"host"`
	Port    string `toml:"port"`
}

func New() *App {
	appconfig := &AppConfig{}
	conf.ReadAppConfFile("app.toml", appconfig)
	webServer := &WebServer{
		Router: httprouter.New(),
	}
	engine := &App{
		Config: appconfig,
		Server: webServer,
	}
	return engine
}

func Test() *httprouter.Router {
	return httprouter.New()
}

type RouterRegister func(*WebServer)

func (app *App) WebServer() *WebServer {
	return app.Server
}

func (app *App) RegisterRouter(rr RouterRegister) {
	rr(app.WebServer())
}

func initRuntimeLog() {
	file, err := os.OpenFile("log/runtime.log",
		os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	log.SetOutput(file)
}

func (app *App) StartServer() {
	initRuntimeLog()
	addr := app.Config.Host + ":" + app.Config.Port
	handler := app.WebServer().Router
	log.Println(addr)
	graceful.StartServer(addr, handler)
}
