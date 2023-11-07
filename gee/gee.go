package gee

import (
	"log"
	"net/http"
)

// HandlerFunc 是一个自定义类型。
// 具体来说，HandlerFunc 是一个函数类型，它接受两个参数，分别是 http.ResponseWriter 和 *http.Request。
// 主要用于定义 HTTP 请求处理函数的格式。在许多 Web 开发框架中，HTTP 请求处理函数通常需要遵循特定的签名，
// 即接受一个 http.ResponseWriter 用于写入 HTTP 响应，以及一个 *http.Request 用于表示 HTTP 请求。
// 通过定义 HandlerFunc 类型，可以使代码更加模块化和可扩展，允许你在不修改函数签名的情况下更改或替换处理函数。
// 这有助于提高代码的灵活性和可维护性，因为你可以轻松地在不同地方使用相同的函数类型，而无需担心函数签名的不匹配。
type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

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
