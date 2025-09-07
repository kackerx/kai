package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/kackerx/kai/app/domain/errors"
	"github.com/kackerx/kai/app/interfaces/api"
	"github.com/kackerx/kai/configs"
)

func CasbinMiddleware(enforcer *casbin.SyncedEnforcer, skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := configs.C.Casbin
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		p := c.Request.URL.Path
		m := c.Request.Method
		if b, err := enforcer.Enforce(api.GetUserID(c), p, m); err != nil {
			api.ResError(c, errors.WithStack(err))
			return
		} else if !b {
			api.ResError(c, errors.ErrNoPerm)
			return
		}
		c.Next()
	}
}
