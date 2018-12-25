package realgo

import "github.com/julienschmidt/httprouter"

type Engine struct {
	ConfPath string
	Router *httprouter.Router


}

func New() *Engine{
	engine := &Engine{}
	return engine
}


