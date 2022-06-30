package gee

import (
	"net/http"
	"strings"
)

type Router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func ParsePattern(pattern string) []string {
	vs := strings.Split(pattern, "/")

	parts := make([]string, 0)
	for _, item := range vs {
		if item != "" {
			parts = append(parts, item)
			if item[0] == '*' {
				break
			}
		}
	}
	return parts
}
func NewRouter() *Router {
	return &Router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	parts := ParsePattern(pattern)

	key := method + "-" + pattern
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler
}

func (r *Router) GetRoute(method string, path string) (*node, map[string]string) {
	searchParts := ParsePattern(path)
	params := make(map[string]string)
	root, ok := r.roots[method]

	if !ok {
		return nil, nil
	}

	n := root.search(searchParts, 0)

	if n != nil {
		parts := ParsePattern(n.Pattern)
		for index, part := range parts {
			if part[0] == ':' {
				params[part[1:]] = searchParts[index]
			}
			if part[0] == '*' && len(part) > 1 {
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}
		}
		return n, params
	}

	return nil, nil
}

func (r *Router) handle(c *Context) {
	n, params := r.GetRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.Pattern
		r.handlers[key](c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FUND:%s\n", c.Path)
	}
}
