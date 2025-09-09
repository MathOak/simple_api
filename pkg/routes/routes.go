package routes

import (
	filmRoutes "simple_api/pkg/models/films/routes"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// importe outros módulos de rotas aqui
)

func RegisterRoutes(r *gin.Engine) {
	// Rota Swagger na rota /docs
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Subrotas
	filmRoutes.RegisterFilmRoutes(r)
	// chame outras funções de registro de subrotas aqui
}
