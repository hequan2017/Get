package main

import (
	"github.com/hequan2017/get/get"
	"net/http"
)

func main() {
	r := get.Default()
	r.GET("/", func(c *get.Context) {
		c.String(http.StatusOK, "Hello getktutu\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *get.Context) {
		names := []string{"getktutu"}
		c.String(http.StatusOK, names[100])
	})

	_ = r.Run(":9999")
}
