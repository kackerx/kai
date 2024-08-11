package kai

import (
	"net/http"
)

type HandleFunc func(ctx *Context)

type Server interface {
	http.Handler
}

type HTTPServer struct {
	*router
	mlds []Middleware
}

type HTTPServerOption func(server *HTTPServer)

func WithMiddleware(mlds ...Middleware) HTTPServerOption {
	return func(server *HTTPServer) {
		server.mlds = mlds
	}
}

func NewHTTPServer(opts ...HTTPServerOption) *HTTPServer {
	server := &HTTPServer{router: newRouter()}
	for _, opt := range opts {
		opt(server)
	}

	return server
}

func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		resp: writer,
	}

	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	node, ok := h.router.findRoute(ctx.Req.Method, ctx.Req.URL.Path)
	if !ok || node == nil {
		ctx.resp.WriteHeader(http.StatusNotFound)
		return
	}

	ctx.pathParams = node.pathParams
	ctx.Route = node.route

	root := node.handler
	for i := len(h.mlds) - 1; i >= 0; i-- {
		root = h.mlds[i](root)
	}
	root(ctx)
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

func (h *HTTPServer) Start(addr string) error {
	return http.ListenAndServe(addr, h)
}
