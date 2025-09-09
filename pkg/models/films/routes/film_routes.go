package routes

import (
	"simple_api/pkg/models/films/services"

	"github.com/gin-gonic/gin"
)

func RegisterFilmRoutes(r *gin.Engine) {
	films := r.Group("/films")
	{
		films.GET("/", func(c *gin.Context) {

			var films, err = services.ListAllFilms(c)

			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, films)
		})
		// Rota de busca com parâmetros
		films.GET("/search", func(c *gin.Context) {
			// Exemplos de parâmetros: title, year, director, order_by, order_dir
			title := c.Query("title")
			year := c.Query("year")
			director := c.Query("director")
			orderBy := c.DefaultQuery("order_by", "")
			orderDir := c.DefaultQuery("order_dir", "asc")

			var films, err = services.SearchFilmsOrdered(c, title, year, director, orderBy, orderDir)
			if err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			c.JSON(200, films)
		})
	}
}
