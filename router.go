package kai

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
		root.route = path
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
	root.route = path
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
			if root.paramChild != nil {
				root = root.paramChild
				root.pathParams[root.path[1:]] = seg
			} else if root.starChildren != nil {
				root = root.starChildren
			} else {
				return nil, false
			}
		} else {
			root = child
		}
	}

	return root, true
}

// ----- node

type node struct {
	path         string
	route        string
	chidren      map[string]*node
	paramChild   *node
	pathParams   map[string]string
	starChildren *node
	handler      HandleFunc
}

func newNode(path string, handler HandleFunc) *node {
	return &node{path: path, chidren: make(map[string]*node), handler: handler, pathParams: make(map[string]string)}
}

func (n *node) childOrCreate(seg string) (child *node) {
	if seg == "*" { // 通配符匹配
		n.starChildren = newNode("*", nil)
		return n.starChildren
	}
	if strings.HasPrefix(seg, ":") { // 参数匹配
		n.paramChild = newNode(seg, nil)
		return n.paramChild
	}
	if child, ok := n.chidren[seg]; ok {
		return child
	}

	child = newNode(seg, nil)
	n.chidren[seg] = child

	return
}
