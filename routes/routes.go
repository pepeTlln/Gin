package routes

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type Usuario struct {
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}

var usuarios []Usuario

func SetupRoutes(r *gin.Engine) {
	r.Static("/static", "./static")
	r.LoadHTMLGlob("templates/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/:page", func(c *gin.Context) {
		page := c.Param("page")
		// Verificar si la página ya tiene la extensión ".html"
		if !strings.HasSuffix(page, ".html") {
			page += ".html" // Agregar la extensión ".html" si no está presente
		}
		// Comprobar si la plantilla solicitada existe
		if _, err := os.Stat("templates/" + page); err == nil {
			// Si la plantilla existe, renderizarla
			c.HTML(http.StatusOK, page, nil)
		} else {
			// Si la plantilla no existe, mostrar página de error 404
			c.HTML(http.StatusNotFound, "404.html", nil)
		}
	})

}
