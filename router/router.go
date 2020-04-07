package router

import (
	"strings"
)

type Router struct {
	nodes map[string]*Node
}

func (router *Router) AddRoute(method string, pattern string) {
	_, ok := router.nodes[method]
	if !ok {
		router.nodes[method] = &Node{}
	}
	router.nodes[method].insert(pattern)
}

func (router *Router) GetRoute(method string, path string) (string, map[string]string) {
	searchParts := parsePattern(path)
	params := make(map[string]string)
	node, ok := router.nodes[method]
	if !ok {
		return "", nil
	}
	node = node.search(path)
	if node != nil {
		for index, part := range parsePattern(node.pattern) {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return node.pattern, params
	}
	return "", nil
}

func New() *Router {
	return &Router{
		nodes: make(map[string]*Node),
	}
}
