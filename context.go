package realgo

import (
	"github.com/greatming/realgo/lib/logger"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type WebContext struct {
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
	Keys     map[string]interface{}
	Logger   *logger.Logger
}

func (ctx *WebContext) Init() {
	ctx.Request.ParseForm()
}

func (ctx *WebContext) Set(key string, value interface{}) {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	ctx.Keys[key] = value
}

func (ctx *WebContext) Get(key string) (value interface{}, exists bool) {
	value, exists = ctx.Keys[key]
	return
}

func (ctx *WebContext) GetInt(key string) (v int) {
	if value, ok := ctx.Keys[key]; ok && value != nil {
		v, _ = value.(int)
	}
	return
}

func (ctx *WebContext) GetInt64(key string) (v int64) {
	if value, ok := ctx.Keys[key]; ok && value != nil {
		v, _ = value.(int64)
	}
	return
}

func (ctx *WebContext) GetString(key string) (v string) {
	if value, ok := ctx.Keys[key]; ok && value != nil {
		v, _ = value.(string)
	}
	return
}
