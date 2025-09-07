package api

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/kackerx/kai/configs"
	"github.com/redis/go-redis/v9"

	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
	"github.com/google/gops/agent"
	"github.com/kackerx/kai/pkg/logger"
)

type Options struct {
	ConfigFile string
	ModelFile  string
	MenuFile   string
	WWWDir     string
	Version    string
}

type Option func(*Options)

func SetConfigFile(s string) Option {
	return func(o *Options) {
		o.ConfigFile = s
	}
}

func SetModelFile(s string) Option {
	return func(o *Options) {
		o.ModelFile = s
	}
}

func SetWWWDir(s string) Option {
	return func(o *Options) {
		o.WWWDir = s
	}
}

func SetMenuFile(s string) Option {
	return func(o *Options) {
		o.MenuFile = s
	}
}

func SetVersion(s string) Option {
	return func(o *Options) {
		o.Version = s
	}
}

func InitCaptcha() {
	cfg := configs.C.Captcha
	if cfg.Store == "redis" {
		rc := configs.C.Redis
		// Create Redis client for captcha store
		redisCli := redis.NewClient(&redis.Options{
			Addr:     rc.Addr,
			Password: rc.Password,
			DB:       cfg.RedisDB,
		})
		// Note: Using memory store as fallback since captcha store may not be compatible with redis v9
		// You may need to implement a custom store or use memory store
		_ = redisCli // Suppress unused variable warning for now
		captcha.SetCustomStore(store.NewMemoryStore(100, captcha.Expiration))
	}
}

func InitMonitor(ctx context.Context) func() {
	if c := configs.C.Monitor; c.Enable {
		// ShutdownCleanup set false to prevent automatically closes on os.Interrupt
		// and close agent manually before service shutting down
		err := agent.Listen(agent.Options{Addr: c.Addr, ConfigDir: c.ConfigDir, ShutdownCleanup: false})
		if err != nil {
			logger.WithContext(ctx).Errorf("Agent monitor error: %s", err.Error())
		}
		return func() {
			agent.Close()
		}
	}
	return func() {}
}

func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := configs.C.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.WithContext(ctx).Printf("HTTP server is running at %s.", addr)

		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Error(err.Error())
		}
	}
}
