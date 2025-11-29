package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/kackerx/kai/app/domain/auth"
	"github.com/kackerx/kai/app/interfaces/api/handler"
)

type Router interface {
	Register(app *gin.Engine) error
	Prefixes() []string
}

func NewRouter(
	auth auth.Repository,
	casbinEnforcer *casbin.SyncedEnforcer,
	loginHandler handler.Login,
	menuHandler handler.Menu,
	roleHandler handler.Role,
	userHandler handler.User,
	healthHandler handler.HealthCheck,
	testHandler handler.Test,
) Router {
	return &router{
		auth:           auth,
		casbinEnforcer: casbinEnforcer,
		loginHandler:   loginHandler,
		menuHandler:    menuHandler,
		roleHandler:    roleHandler,
		userHandler:    userHandler,
		healthHandler:  healthHandler,
		testHandler:    testHandler,
	}
}

type router struct {
	auth           auth.Repository
	casbinEnforcer *casbin.SyncedEnforcer
	loginHandler   handler.Login
	menuHandler    handler.Menu
	roleHandler    handler.Role
	userHandler    handler.User
	healthHandler  handler.HealthCheck
	testHandler    handler.Test
}

func (a *router) Register(app *gin.Engine) error {
	a.RegisterAPI(app)
	return nil
}

func (a *router) Prefixes() []string {
	return []string{
		"/api/",
	}
}
