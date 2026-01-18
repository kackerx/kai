//go:build wireinject
// +build wireinject

package test

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/kackerx/kai/app/application"
	"github.com/kackerx/kai/app/infrastructure/article"
	"github.com/kackerx/kai/app/interfaces/api/handler"
)

var handlerSet = wire.NewSet(
	handler.NewArticle,
	handler.NewHealthCheck,
	handler.NewLogin,
	handler.NewMenu,
	handler.NewRole,
	handler.NewUser,
	handler.NewTest,
)

var appSet = wire.NewSet(
	application.NewArticle,
)

var repositorySet = wire.NewSet(
	article.NewRepository,
)

func BuildEngine() (*gin.Engine, func(), error) {
	wire.Build(
		NewDB,
		InitGinEngine,
		InitRouter,

		handlerSet,
		appSet,
		repositorySet,
	)
	return nil, nil, nil
}
