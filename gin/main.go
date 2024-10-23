package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	g := gin.Default()

	group := g.Group("/user", func(c *gin.Context) {

	})

	// group.GET()

	fmt.Println(group)
}
