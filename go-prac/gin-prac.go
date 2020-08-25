package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Conan")
	})
	fmt.Println("listen and serve on http://0.0.0.0:8080")
	r.Run() // listen and serve on 0.0.0.0:8080
}