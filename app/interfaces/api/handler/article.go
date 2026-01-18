package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kackerx/kai/app/application"
	"github.com/kackerx/kai/app/interfaces/api"
	"github.com/kackerx/kai/app/interfaces/api/request"
)

type Article interface {
	Edit(c *gin.Context)
}

func NewArticle(articleApp application.Article) Article {
	return &articleHandler{
		app: articleApp,
	}
}

type articleHandler struct {
	app application.Article
}

func (a *articleHandler) Edit(c *gin.Context) {
	ctx := c.Request.Context()
	var req request.EditArticleReq
	if err := api.ParseJSON(c, &req); err != nil {
		api.ResError(c, err)
		return
	}

	if resp, err := a.app.Edit(ctx, &req); err != nil {
		api.ResError(c, err)
	} else {
		api.ResSuccess(c, resp)
	}
}
