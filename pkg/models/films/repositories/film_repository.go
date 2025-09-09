package repositories

import (
	"simple_api/pkg/models/films/dto"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type FilmRepository struct {
	DB *gorm.DB
}

func GetAllFilms(c *gin.Context) ([]dto.Films, error) {
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "DB connection not found"})
		return nil, nil
	}

	var films []dto.Films
	if err := db.Find(&films).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return nil, err
	}

	return films, nil
}

func (repo *FilmRepository) GetFilm(c *gin.Context, id string) (dto.FilmDTO, error) {
	// Implementação de consulta ao banco
	return dto.FilmDTO{}, nil
}

func (repo *FilmRepository) SaveFilm(c *gin.Context, film dto.FilmDTO) error {
	// Implementação de salvamento no banco
	return nil
}

func (repo *FilmRepository) UpdateFilm(c *gin.Context, id string, film dto.FilmDTO) error {
	// Implementação de atualização no banco
	return nil
}
