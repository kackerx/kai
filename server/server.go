package server

import "net/http"

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
}

type HTTPServer struct {
	*router
}

func NewHTTPServer() *HTTPServer {
	return &HTTPServer{router: newRouter()}
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		req:  request,
		resp: writer,
	}

	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	node, ok := h.router.findRoute(ctx.req.Method, ctx.req.URL.Path)
	if !ok || node == nil {
		ctx.resp.WriteHeader(http.StatusNotFound)
		return
	}

	node.handler(ctx)
}

func (h *HTTPServer) AddRoute(method, path string, handler HandleFunc) {
	h.router.addRoute(method, path, handler)
}

func (h *HTTPServer) Get(path string, handler HandleFunc) {
	h.router.addRoute(http.MethodGet, path, handler)
}

func (h *HTTPServer) Post(path string, handler HandleFunc) {
	h.router.addRoute(http.MethodPost, path, handler)
}
