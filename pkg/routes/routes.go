package routes

import (
	filmRoutes "simple_api/pkg/models/films/routes"

	"github.com/gin-gonic/gin"
	// importe outros módulos de rotas aqui
)

func RegisterRoutes(r *gin.Engine) {
	// Rotas principais
	r.GET("/docs", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Documentação"})
	})

	// Subrotas
	filmRoutes.RegisterFilmRoutes(r)
	// chame outras funções de registro de subrotas aqui
}
