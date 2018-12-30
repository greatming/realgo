package realgo

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

type WebContext struct {
	Response http.ResponseWriter
	Request *http.Request
	Params httprouter.Params
	keys map[string]interface{}
}

func (ctx *WebContext)Init()  {
	ctx.Request.ParseForm()
}

func (ctx *WebContext)Set(key string, value interface{})  {
	if ctx.keys == nil{
		ctx.keys = make(map[string]interface{})
	}
	ctx.keys[key] = value
}

func (ctx *WebContext)Get(key string)(value interface{}, exists bool)  {
	value, exists = ctx.keys[key]
	return
}

func (ctx *WebContext)GetInt(key string)(v int)  {
	if value, ok := ctx.keys[key]; ok && value != nil{
		v, _ = value.(int)
	}
	return
}

func (ctx *WebContext)GetInt64(key string)(v int64)  {
	if value, ok := ctx.keys[key]; ok && value != nil{
		v, _ = value.(int64)
	}
	return
}

func (ctx *WebContext)GetString(key string)(v string)  {
	if value, ok := ctx.keys[key]; ok && value != nil{
		v, _ = value.(string)
	}
	return
}

