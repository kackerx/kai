package accesslog

import (
	"fmt"
	"net/http"
	"testing"

	"kai"
)

func TestAccessLog(t *testing.T) {
	builder := new(MiddlewareBuilder)
	accessMiddleware := builder.BuildLogFunc(func(s string) {
		fmt.Println("access_log: ", s)
	}).Build()

	server := kai.NewHTTPServer(kai.WithMiddleware(accessMiddleware))

	server.Post("/a/b/*", func(c *kai.Context) {
		c.Html(200, "hello world")
	})
	request, _ := http.NewRequest(http.MethodPost, "/a/b/c", nil)
	server.ServeHTTP(nil, request)
}
