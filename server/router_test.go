package server

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_router_AddRoute(t *testing.T) {
	type fields struct {
		trees map[string]*node
	}
	type args struct {
		method  string
		path    string
		handler HandleFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "/",
			fields: fields{map[string]*node{}},
			args: args{
				method: http.MethodGet,
				path:   "/",
				handler: func(ctx *Context) {
					fmt.Println("/")
				},
			},
		},
		{
			name:   "user",
			fields: fields{map[string]*node{}},
			args: args{
				method: http.MethodGet,
				path:   "/user/info",
				handler: func(ctx *Context) {
					fmt.Println("/")
				},
			},
		},
		{
			name:   "hehe",
			fields: fields{map[string]*node{}},
			args: args{
				method: http.MethodGet,
				path:   "/user/detail/110",
				handler: func(ctx *Context) {
					fmt.Println("hehe")
				},
			},
		},
		{
			name:   "order",
			fields: fields{map[string]*node{}},
			args: args{
				method: http.MethodGet,
				path:   "/order/detail/110",
				handler: func(ctx *Context) {
					fmt.Println("hehe")
				},
			},
		},
		{
			name:   "a",
			fields: fields{map[string]*node{}},
			args: args{
				method: http.MethodPost,
				path:   "/order/*",
				handler: func(ctx *Context) {
					fmt.Println("hehe")
				},
			},
		},
		{
			name:   "asdf",
			fields: fields{map[string]*node{}},
			args: args{
				method: http.MethodPost,
				path:   "/order/add",
				handler: func(ctx *Context) {
					fmt.Println("hehe")
				},
			},
		},
	}
	r := &router{
		trees: map[string]*node{},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			r.addRoute(tt.args.method, tt.args.path, tt.args.handler)
		})
	}
	fmt.Println(len(r.trees))
}

func Test_router_findRoute(t *testing.T) {
	testRoute := []struct {
		method  string
		path    string
		handler HandleFunc
	}{
		{
			method: http.MethodGet,
			path:   "/user/detail/110",
			handler: func(ctx *Context) {
				fmt.Println("hehe")
			},
		},
		{
			method: http.MethodPost,
			path:   "/order/add",
			handler: func(ctx *Context) {
				fmt.Println("hehe")
			},
		},
		{
			method: http.MethodPost,
			path:   "/order/*",
			handler: func(ctx *Context) {
				fmt.Println("hehe")
			},
		},
	}
	r := newRouter()
	for _, route := range testRoute {
		r.addRoute(route.method, route.path, route.handler)
	}
	// ---
	tests := []struct {
		name string

		method string
		path   string

		wantRes *node
		wantOk  bool
	}{
		{

			"hehe",
			http.MethodGet,
			"/user/detail/110",
			&node{
				path:    "110",
				chidren: nil,
				handler: nil,
			},
			true,
		},
		{
			"asdf",
			http.MethodPost,
			"/order/add",
			&node{
				path:    "add",
				chidren: nil,
				handler: nil,
			},
			true,
		},
		{
			"asf",
			http.MethodPost,
			"/order/abc",
			&node{
				path:    "*",
				chidren: nil,
				handler: nil,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRes, gotOk := r.findRoute(tt.method, tt.path)
			if !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("findRoute() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
			if gotOk != tt.wantOk {
				t.Errorf("findRoute() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestFoo(t *testing.T) {
	mockHandler := func(ctx Context) {
		fmt.Println("mockHandler")
	}
	a := mockHandler
	b := mockHandler

	af := reflect.ValueOf(a)
	bf := reflect.ValueOf(b)

	assert.Equal(t, af, bf)
}
