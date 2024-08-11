package kai

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Req        *http.Request
	resp       http.ResponseWriter
	Route      string
	pathParams map[string]string
}

// --- Input
func (ctx *Context) BindJSON(obj any) error {
	return json.NewDecoder(ctx.Req.Body).Decode(obj)
}

func (ctx *Context) Form(key string) (string, error) {
	if err := ctx.Req.ParseForm(); err != nil {
		return "", err
	}

	return ctx.Req.FormValue(key), nil
}

func (ctx *Context) Query(key string) (string, error) {
	return ctx.Req.URL.Query().Get(key), nil
}

func (ctx *Context) Path(key string) (string, error) {
	return ctx.pathParams[key], nil
}

func (ctx *Context) Html(code int, html string) {
	ctx.resp.WriteHeader(code)
	ctx.resp.Header().Set("Content-Type", "text/html")
	ctx.resp.Write([]byte(html))
}

func (ctx *Context) Url() string {
	return ctx.Req.URL.Path
}

func (ctx *Context) Get(key string) string {
	return ctx.pathParams[key]
}

func (ctx *Context) SetHeader(headerMap map[string]string) {
	for key, value := range headerMap {
		ctx.resp.Header().Add(key, value)
	}
}

func (ctx *Context) GetHeader(key string) []string {
	return ctx.Req.Header[key]
}
