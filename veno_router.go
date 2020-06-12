package veno

import (
	"github.com/qq15725/veno/router"
	"net/http"
)

type Router struct {
	*router.Router
	handlers map[string]Handler
}

func (router *Router) GET(pattern string, handler Handler) {
	router.AddRoute("GET", pattern)
	router.handlers["GET-"+pattern] = handler
}

func (router *Router) POST(pattern string, handler Handler) {
	router.AddRoute("POST", pattern)
	router.handlers["POST-"+pattern] = handler
}

func (router *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	pattern, params := router.GetRoute(req.Method, req.URL.Path)
	ctx := newHttpContext(res, req)
	if pattern != "" {
		ctx.Params = params
		router.handlers[req.Method+"-"+pattern](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}

func newRouter() *Router {
	return &Router{
		Router:   router.New(),
		handlers: make(map[string]Handler),
	}
}
