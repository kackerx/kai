package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/kackerx/kai/app/domain/contextx"
	"github.com/kackerx/kai/pkg/logger"
	"github.com/kackerx/kai/pkg/util/trace"
)

func TraceMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		traceID := c.GetHeader("X-Request-Id")
		if traceID == "" {
			traceID = trace.NewTraceID()
		}

		ctx := contextx.NewTraceID(c.Request.Context(), traceID)
		ctx = logger.NewTraceIDContext(ctx, traceID)
		c.Request = c.Request.WithContext(ctx)
		c.Writer.Header().Set("X-Trace-Id", traceID)

		c.Next()
	}
}
