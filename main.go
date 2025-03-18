package main

import (
	"gin/routes"

	"github.com/gin-gonic/gin"
)

// go get -u github.com/gin-gonic/gin
// TB se podría go install github.com/gin-gonic/gin@latest
func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
