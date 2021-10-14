package gee

import "net/http"

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine{
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string,pattern string, handlerFunc HandlerFunc){
	key:= method +"-"+ pattern
	engine.router[key] = handlerFunc
}
