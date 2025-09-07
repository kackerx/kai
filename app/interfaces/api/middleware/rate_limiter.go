package middleware

import (
	"strconv"
	"time"

	"github.com/kackerx/kai/app/domain/errors"

	"github.com/kackerx/kai/app/interfaces/api"

	"github.com/kackerx/kai/configs"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v10"
	"github.com/redis/go-redis/v9"
)

func RateLimiterMiddleware(skippers ...SkipperFunc) gin.HandlerFunc {
	cfg := configs.C.RateLimiter
	if !cfg.Enable {
		return EmptyMiddleware()
	}

	rc := configs.C.Redis
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": rc.Addr,
		},
		Password: rc.Password,
		DB:       cfg.RedisDB,
	})

	limiter := redis_rate.NewLimiter(ring)

	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}

		userID := api.GetUserID(c)
		if userID != "" {
			limit := cfg.Count
			ctx := c.Request.Context()
			res, err := limiter.Allow(ctx, userID, redis_rate.PerMinute(int(limit)))
			if err != nil {
				api.ResError(c, errors.ErrInternalServer)
				return
			}
			allowed := res.Allowed > 0
			rate := res.Allowed
			delay := res.RetryAfter
			if !allowed {
				h := c.Writer.Header()
				h.Set("X-RateLimit-Limit", strconv.FormatInt(limit, 10))
				h.Set("X-RateLimit-Remaining", strconv.FormatInt(limit-int64(rate), 10))
				delaySec := int64(delay / time.Second)
				h.Set("X-RateLimit-Delay", strconv.FormatInt(delaySec, 10))
				api.ResError(c, errors.ErrTooManyRequests)
				return
			}
		}

		c.Next()
	}
}
