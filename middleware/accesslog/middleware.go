package accesslog

import (
	jsoniter "github.com/json-iterator/go"

	"kai"
)

type MiddlewareBuilder struct {
	logFunc func(string)
}

func (m *MiddlewareBuilder) BuildLogFunc(fn func(string)) *MiddlewareBuilder {
	m.logFunc = fn
	return m
}

func (m *MiddlewareBuilder) Build() kai.Middleware {
	return func(next kai.HandleFunc) kai.HandleFunc {
		return func(ctx *kai.Context) {
			// 记录请求
			defer func() {
				log := newAccessLog(ctx.Req.Host, ctx.Route, ctx.Req.URL.Path, ctx.Req.Method)
				logString, _ := jsoniter.MarshalToString(log)
				m.logFunc(logString)
			}()

			next(ctx)
		}
	}
}

type accessLog struct {
	Host       string `json:"host,omitempty"`
	Route      string `json:"route,omitempty"`
	Path       string `json:"path,omitempty"`
	HTTPMethod string `json:"http_method,omitempty"`
}

func newAccessLog(host string, route string, path string, HTTPMethod string) accessLog {
	return accessLog{Host: host, Route: route, Path: path, HTTPMethod: HTTPMethod}
}
