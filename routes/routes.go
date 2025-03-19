package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	r.GET("/usuario/:nombre", func(c *gin.Context) { // :nombre es un parámetro
		nombre := c.Param("nombre")
		c.String(http.StatusOK, "Hola %s", nombre)
	})
	r.POST("/usuarios", func(c *gin.Context) {
		var usuario Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if usuario.Nombre == "" || usuario.Email == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "nombre y email son requeridos"})
			return
		}
		usuarios = append(usuarios, usuario)
		c.JSON(http.StatusCreated, usuarios)

	})
}
