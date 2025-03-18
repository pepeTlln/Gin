package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/usuario/:nombre", func(c *gin.Context) { // :nombre es un parámetro
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola %s", nombre)
	})

}
