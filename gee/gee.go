package gee

import (
	"fmt"
	"net/http"
)

// HandlerFunc 创建http.handlerfunc
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine  创建一个router结构体
type Engine struct {
	router map[string]HandlerFunc
}

// New 新建一个router
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// addRoute 添加一个route :method 传递方法，pattern 模式,handler 方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 创建一个GET方法,传递参数 :pattern 模式， handler 方法
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 创建一个POST方法,传递参数 :pattern 模式， handler 方法
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
