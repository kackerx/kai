//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/kackerx/kai/app/application"
	"github.com/kackerx/kai/app/domain/menu"
	"github.com/kackerx/kai/app/domain/user"
	menuInfra "github.com/kackerx/kai/app/infrastructure/menu"
	menuActionInfra "github.com/kackerx/kai/app/infrastructure/menu/menuaction"
	menuActionResourceInfra "github.com/kackerx/kai/app/infrastructure/menu/menuactionresource"

	rbacInfra "github.com/kackerx/kai/app/infrastructure/rbac"
	transInfra "github.com/kackerx/kai/app/infrastructure/trans"
	userInfra "github.com/kackerx/kai/app/infrastructure/user"
	roleInfra "github.com/kackerx/kai/app/infrastructure/user/role"
	roleMenuInfra "github.com/kackerx/kai/app/infrastructure/user/rolemenu"
	userRoleInfra "github.com/kackerx/kai/app/infrastructure/user/userrole"
	"github.com/kackerx/kai/app/interfaces/api/handler"
	"github.com/kackerx/kai/app/interfaces/api/router"
	"github.com/kackerx/kai/injector/api"

	"github.com/google/wire"
)

func BuildApiInjector() (*ApiInjector, func(), error) {
	wire.Build(
		// init,
		InitGormDB,
		api.InitAuth,
		api.InitGinEngine,
		api.InitCasbin,

		// domain
		user.NewService,
		menu.NewService,

		// infrastructure
		menuInfra.NewRepository,
		menuActionInfra.NewRepository,
		menuActionResourceInfra.NewRepository,
		userInfra.NewRepository,
		userRoleInfra.NewRepository,
		roleMenuInfra.NewRepository,
		roleInfra.NewRepository,
		transInfra.NewRepository,
		rbacInfra.NewRepository,

		// application
		application.NewMenu,
		application.NewRole,
		application.NewUser,
		application.NewLogin,
		application.NewRbacAdapter,
		application.NewSeed,

		// handler
		handler.NewHealthCheck,
		handler.NewUser,
		handler.NewRole,
		handler.NewMenu,
		handler.NewLogin,

		// router
		router.NewRouter,

		// injector
		NewApiInjector,
	)
	return nil, nil, nil
}
