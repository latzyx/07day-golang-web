package gee

import (
	"net/http"
)

// HandlerFunc 创建http.handlerfunc
// "HandlerFunc" defines the request handler used by gee
type HandlerFunc func(*Context)

// Engine  创建一个router结构体 implement the interface of ServeHTTP
type Engine struct {
	router *router
}

// New 新建一个router
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute 添加一个route :method 传递方法，pattern 模式,handler 方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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
	c := newContext(w, req)
	engine.router.handle(c)
}
