package services

import (
	"simple_api/pkg/models/films/dto"
	"simple_api/pkg/models/films/repositories"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FilmService struct {
	Repo *repositories.FilmRepository
}

func ListAllFilms(c *gin.Context) ([]dto.FilmDTO, error) {
	films, err := repositories.GetAllFilms(c)
	if err != nil {
		return nil, err
	}
	var filmsDTO []dto.FilmDTO
	for _, film := range films {
		filmsDTO = append(filmsDTO, dto.FilmDTO{
			ID:       film.ID,
			Title:    film.Title,
			Year:     film.Year,
			Genre:    film.Genre,
			Director: film.Director,
			Link:     film.Link,
		})
	}
	return filmsDTO, nil
}

// Busca filmes por parâmetros (exemplo: title, year, director)
func SearchFilms(c *gin.Context, title, year, director string) ([]dto.FilmDTO, error) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "DB connection not found"})
		return nil, nil
	}

	var films []dto.Films
	query := db.Model(&dto.Films{})
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if year != "" {
		query = query.Where("year = ?", year)
	}
	if director != "" {
		query = query.Where("director LIKE ?", "%"+director+"%")
	}
	if err := query.Find(&films).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return nil, err
	}

	var filmsDTO []dto.FilmDTO
	for _, film := range films {
		filmsDTO = append(filmsDTO, dto.FilmDTO{
			ID:       film.ID,
			Title:    film.Title,
			Year:     film.Year,
			Genre:    film.Genre,
			Director: film.Director,
			Link:     film.Link,
		})
	}
	return filmsDTO, nil
}

// Busca filmes por parâmetros e ordenação
func SearchFilmsOrdered(c *gin.Context, title, year, director, orderBy, orderDir string) ([]dto.FilmDTO, error) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "DB connection not found"})
		return nil, nil
	}

	var films []dto.Films
	query := db.Model(&dto.Films{})
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}
	if year != "" {
		query = query.Where("year = ?", year)
	}
	if director != "" {
		query = query.Where("director LIKE ?", "%"+director+"%")
	}
	// Ordenação
	if orderBy != "" {
		dir := "asc"
		if orderDir == "desc" {
			dir = "desc"
		}
		query = query.Order(orderBy + " " + dir)
	}
	if err := query.Find(&films).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return nil, err
	}

	var filmsDTO []dto.FilmDTO
	for _, film := range films {
		filmsDTO = append(filmsDTO, dto.FilmDTO{
			ID:       film.ID,
			Title:    film.Title,
			Year:     film.Year,
			Genre:    film.Genre,
			Director: film.Director,
			Link:     film.Link,
		})
	}
	return filmsDTO, nil
}
