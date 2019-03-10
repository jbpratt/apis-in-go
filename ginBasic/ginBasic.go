package main

import (
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// GET takes a route and handler func.
	// Handler takes the gin context obj
	r.GET("/pingTime", func(c *gin.Context) {
		// json serializer available on gin context
		c.JSON(200, gin.H{
			"serverTime": time.Now().UTC(),
		})
	})
	r.Run(":8000") // Listen and server on localhost:8000
}
