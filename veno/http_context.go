package veno

import (
	"github.com/qq15725/go/context"
	"net/http"
)

type HttpContext struct {
	*context.HttpContext
}

func newHttpContext(res http.ResponseWriter, req *http.Request) *HttpContext {
	return &HttpContext{
		HttpContext: context.NewHttpContext(res, req),
	}
}
