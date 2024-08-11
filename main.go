package kai

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name" validate:"required"`
	Age  int    `json:"age"	`
}

func main() {
	gin.Default()
	kai := NewHTTPServer()

	kai.AddRoute(http.MethodGet, "/hehe", func(ctx *Context) {
		ctx.Html(http.StatusOK, "<h1>Hello Kai</h1>")
	})

	kai.Post("/order/detail", func(c *Context) {
		var u User
		if err := c.BindJSON(&u); err != nil {
			c.Html(http.StatusBadRequest, "bad request")
			return
		}

		fmt.Println(u)
		c.Html(http.StatusOK, "detail")
	})

	kai.Get("/order/*", func(c *Context) {
		c.Html(http.StatusOK, c.Url())
	})

	kai.Get("/order/detail/:id", func(c *Context) {
		c.Html(http.StatusOK, c.Get("id"))
	})

	fmt.Println("listen on [:9999]")
	if err := http.ListenAndServe(":9999", kai); err != nil {
		panic(err)
	}
}

func foo() {

}
