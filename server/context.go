package server

import (
	"net/http"
)

type Context struct {
	req        *http.Request
	resp       http.ResponseWriter
	pathParams map[string]string
}

func (ctx *Context) Html(code int, html string) {
	ctx.resp.WriteHeader(code)
	ctx.resp.Header().Set("Content-Type", "text/html")
	ctx.resp.Write([]byte(html))
}

func (ctx *Context) Path() string {
	return ctx.req.URL.Path
}

func (ctx *Context) Get(key string) string {
	return ctx.pathParams[key]
}
