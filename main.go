package main

import "github.com/gin-gonic/gin"

// go get -u github.com/gin-gonic/gin
// TB se podría go install github.com/gin-gonic/gin@latest
func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8080") // listen and serve on
}
