package main

import (
	"github.com/hequan2017/get/get"
	"net/http"
)

func main() {
	r := get.Default()

	r.GET("/", func(c *get.Context) {
		c.String(http.StatusOK, "Hello get \n")
	})

	r.GET("/get", func(c *get.Context) {
		data := make(map[string]interface{})
		data["name"] = "get"
		c.JSON(http.StatusOK, data)
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", func(c *get.Context) {
			// expect /hello?name=get
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *get.Context) {
			// expect /hello/get
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *get.Context) {
			c.JSON(http.StatusOK, get.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	_ = r.Run(":80")
}
