package realgo

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type WebHandlerFunc func(ctx *WebContext)

type WebServer struct {
	Router *httprouter.Router
	middlewareList []WebHandlerFunc
	endMiddlewareList []WebHandlerFunc
}



func (s *WebServer)GET(path string, handlerFunc WebHandlerFunc)  {
	handler := s.genHandler(handlerFunc)
	s.Router.GET(path, handler)
}

func (s *WebServer)POST(path string, handlerFunc WebHandlerFunc)  {
	handler := s.genHandler(handlerFunc)
	s.Router.POST(path, handler)
}

func (s *WebServer)USE(funcs ...WebHandlerFunc)  {
	for _, f := range funcs{
		s.middlewareList = append(s.middlewareList, f)
	}
}

func (s *WebServer)EUSE(funcs ...WebHandlerFunc)  {
	for _, f := range funcs{
		s.endMiddlewareList = append(s.endMiddlewareList, f)
	}
}


func (s *WebServer)genHandler(handlerFunc WebHandlerFunc) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		context := &WebContext{
			Response:w,
			Request:r,
			Params:p,
		}
		context.Init()
		for _, midFunc := range s.middlewareList{
			midFunc(context)
		}
		handlerFunc(context)
		for _, emidFunc := range s.endMiddlewareList{
			emidFunc(context)
		}
	}
}
