package accesslog

import (
	"fmt"
	"testing"

	"kai"
)

func TestAccessLogE2E(t *testing.T) {
	builder := new(MiddlewareBuilder)
	accessMiddleware := builder.BuildLogFunc(func(s string) {
		fmt.Println("access_log: ", s)
	}).Build()

	server := kai.NewHTTPServer(kai.WithMiddleware(accessMiddleware))

	server.Post("/a/b/*", func(c *kai.Context) {
		c.Html(200, "hello world")
	})

	if err := server.Start(":9999"); err != nil {
		panic(err)
	}
}
