package context

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpContext struct {
	Res        http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
	Params     map[string]string
}

func (ctx *HttpContext) PostForm(key string) string {
	return ctx.Req.FormValue(key)
}

func (ctx *HttpContext) Query(key string) string {
	return ctx.Req.URL.Query().Get(key)
}

func (ctx *HttpContext) Status(code int) {
	ctx.StatusCode = code
	ctx.Res.WriteHeader(code)
}

func (ctx *HttpContext) SetHeader(key string, value string) {
	ctx.Res.Header().Set(key, value)
}

func (ctx *HttpContext) String(code int, format string, values ...interface{}) {
	ctx.SetHeader("Content-Type", "text/plain")
	ctx.Status(code)
	_, err := ctx.Res.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		http.Error(ctx.Res, err.Error(), 500)
	}
}

func (ctx *HttpContext) JSON(code int, obj interface{}) {
	ctx.SetHeader("Content-Type", "application/json")
	ctx.Status(code)
	encoder := json.NewEncoder(ctx.Res)
	if err := encoder.Encode(obj); err != nil {
		http.Error(ctx.Res, err.Error(), 500)
	}
}

func (ctx *HttpContext) Data(code int, data []byte) {
	ctx.Status(code)
	_, err := ctx.Res.Write(data)
	if err != nil {
		http.Error(ctx.Res, err.Error(), 500)
	}
}

func (ctx *HttpContext) HTML(code int, html string) {
	ctx.SetHeader("Content-Type", "text/html")
	ctx.Status(code)
	_, err := ctx.Res.Write([]byte(html))
	if err != nil {
		http.Error(ctx.Res, err.Error(), 500)
	}
}

func NewHttpContext(res http.ResponseWriter, req *http.Request) *HttpContext {
	return &HttpContext{
		Res:    res,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}
