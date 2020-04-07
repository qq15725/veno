package veno

import "github.com/qq15725/go/router"

type Router struct {
	*router.Router
}

func newRouter() *Router {
	return &Router{
		Router: router.New(),
	}
}
