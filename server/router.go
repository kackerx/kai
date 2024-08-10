package server

import (
	"fmt"
	"strings"
)

type router struct {
	trees map[string]*node
}

func newRouter() *router {
	return &router{trees: make(map[string]*node)}
}

func (r *router) addRoute(method string, path string, handler HandleFunc) {
	root, ok := r.trees[method]
	// 根节点不存在，则创建根节点
	if !ok {
		root = newNode("/", nil)
	}

	r.trees[method] = root

	if path == "/" {
		root.handler = handler
		return
	}

	// 切割path
	segs := strings.Split(path, "/")
	for _, seg := range segs[1:] {
		// 递归找位置, 中途节点不存在就新建
		child := root.childOrCreate(seg)
		root = child
	}

	if root.handler != nil {
		fmt.Printf("重复注册路由: %s\n", path)
	}
	root.handler = handler
}

func (r *router) findRoute(method string, path string) (res *node, ok bool) {
	root := r.trees[method]
	if path == "/" {
		return root, true
	}

	segs := strings.Split(path, "/")
	for _, seg := range segs[1:] {
		var child *node
		if child, ok = root.chidren[seg]; !ok {
			return root.starChildren, root.starChildren != nil
		}

		root = child
	}

	return root, true
}

// ----- node

type node struct {
	path         string
	chidren      map[string]*node
	starChildren *node
	handler      HandleFunc
}

func newNode(path string, handler HandleFunc) *node {
	return &node{path: path, chidren: make(map[string]*node), handler: handler}
}

func (n *node) childOrCreate(seg string) (child *node) {
	if seg == "*" {
		n.starChildren = newNode("*", nil)
		return n.starChildren
	}
	if child, ok := n.chidren[seg]; ok {
		return child
	}

	child = newNode(seg, nil)
	n.chidren[seg] = child

	return
}
