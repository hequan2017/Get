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
		c.String(http.StatusOK, "get")
	})

	_ = r.Run(":80")
}
