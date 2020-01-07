package get

import (
	"log"
	"time"
)

func Logger() HandlerFunc {
	return func(c *Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("中间件 Logger [%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
