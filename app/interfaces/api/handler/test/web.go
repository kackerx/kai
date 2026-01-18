package test

import (
	"github.com/gin-gonic/gin"
	"github.com/kackerx/kai/app/interfaces/api/handler"
	"github.com/kackerx/kai/app/interfaces/api/router"
)

func InitGinEngine(r router.Router) *gin.Engine {
	gin.SetMode(gin.TestMode)

	app := gin.New()

	r.Register(app)

	return app
}

type TestRouter struct {
	articleHandler handler.Article
}

func (r *TestRouter) Register(app *gin.Engine) error {
	g := app.Group("/api/v1/articles")
	g.POST("", r.articleHandler.Edit)
	return nil
}

func (r *TestRouter) Prefixes() []string {
	return []string{"/api/v1/articles"}
}

func InitRouter(articleHandler handler.Article) router.Router {
	return &TestRouter{
		articleHandler: articleHandler,
	}
}
