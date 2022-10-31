package router

import (
	"net/http"
)

type Router struct {
	router *PathRoute
}

func ApiRouter() *Router {
	return &Router{
		router: ConfigPath(),
	}
}

func (r *Router) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := r.router.rules[path]

	if !exist {
		r.router.rules[path] = make(map[string]http.HandlerFunc)
	}

	r.router.rules[path][method] = handler
}

func (r *Router) InitRoute() {
	http.Handle("/", r.router)
}
