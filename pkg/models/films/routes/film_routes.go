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
		// @Summary Busca filmes com filtros e ordenação
		// @Description Retorna filmes filtrados por título, ano, diretor e ordenados por campo/direção
		// @Tags films
		// @Accept json
		// @Produce json
		// @Param title query string false "Título do filme"
		// @Param year query int false "Ano do filme"
		// @Param director query string false "Diretor do filme"
		// @Param order_by query string false "Campo para ordenação"
		// @Param order_dir query string false "Direção da ordenação (asc ou desc)"
		// @Success 200 {array} dto.FilmDTO
		// @Failure 500 {object} map[string]string
		// @Router /films/search [get]
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
