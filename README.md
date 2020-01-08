# Go  web framework 

![版本](https://img.shields.io/badge/release-1.0.0-blue.svg)


## 开发者
* 何全


## 特别感谢

> 在此 https://github.com/geektutu/7days-golang 项目上进行修改完成。

> https://geektutu.com/post/gee.html

## 结构

* context 上下文
* router 路由
* middleware 中间件


## context
```go

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index    int
	// engine pointer
	engine *Engine
}

```