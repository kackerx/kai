package main

import (
	"fmt"
	"net/http"

	"kai/server"
)

func main() {
	kai := server.NewHTTPServer()

	kai.AddRoute(http.MethodGet, "/hehe", func(ctx *server.Context) {
		ctx.Html(http.StatusOK, "<h1>Hello Kai</h1>")
	})

	kai.Get("/order/detail", func(ctx *server.Context) {
		ctx.Html(http.StatusOK, "detail")
	})

	kai.Get("/order/*", func(ctx *server.Context) {
		ctx.Html(http.StatusOK, ctx.Path())
	})

	fmt.Println("listen on [:9999]")
	if err := http.ListenAndServe(":9999", kai); err != nil {
		panic(err)
	}
}
