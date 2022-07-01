package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

// Context 创建一个http请求上下文
type Context struct {
	// Writer httpResqonse
	Writer http.ResponseWriter
	// Req http请求
	Req *http.Request
	// Path 路由路径
	Path string
	// Method http请求方式
	Method string
	// Params 路由节点
	Params map[string]string
	// StatusCode http状态码
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// Param 获取路由节点
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// PostForm  Post 表单字段传输方式
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query  GET 查询
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status  http 状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHander  往hander插入数据
func (c *Context) SetHander(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 接收字段值
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHander("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	fmt.Println(values)
}

// JSON  JSON内容解析
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHander("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
	fmt.Printf(" encoder is :%v\n", encoder)
}

// Data  DATA 数据字节
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML  HTML方法
func (c *Context) HTML(code int, html string) {
	c.SetHander("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
