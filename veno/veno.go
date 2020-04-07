package veno

import (
	"net/http"
)

type Application struct {
	router   *Router
	handlers map[string]Handler
}

func (app *Application) GET(pattern string, handler Handler) {
	app.router.AddRoute("GET", pattern)
	app.handlers["GET-"+pattern] = handler
}

func (app *Application) POST(pattern string, handler Handler) {
	app.router.AddRoute("POST", pattern)
	app.handlers["POST-"+pattern] = handler
}

func (app *Application) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	pattern, params := app.router.GetRoute(req.Method, req.URL.Path)
	ctx := newHttpContext(res, req)
	if pattern != "" {
		ctx.Params = params
		app.handlers[req.Method+"-"+pattern](ctx)
	} else {
		ctx.String(http.StatusNotFound, "404 NOT FOUND: %s\n", ctx.Path)
	}
}

func (app *Application) Run(addr string) (err error) {
	return http.ListenAndServe(addr, app)
}

func New() *Application {
	return &Application{
		router:   newRouter(),
		handlers: make(map[string]Handler),
	}
}
