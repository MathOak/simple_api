package routes

import (
	"simple_api/pkg/models/films/repositories"
	"simple_api/pkg/models/films/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterFilmRoutes(r *gin.Engine) {
	films := r.Group("/films")
	{
		films.GET("/", func(c *gin.Context) {
			// Recupera o *gorm.DB do contexto global (defina via middleware se necessário)
			db, ok := c.MustGet("db").(*gorm.DB)
			if !ok {
				c.JSON(500, gin.H{"error": "DB connection not found"})
				return
			}
			repo := &repositories.FilmRepository{DB: db}
			service := &services.FilmService{Repo: repo}

			var films, err = service.ListAllFilms()

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, films)
		})
		// Adicione outras rotas relacionadas a filmes aqui
	}
}
