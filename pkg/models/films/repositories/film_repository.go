package repositories

import (
	"simple_api/pkg/models/films/dto"

	"gorm.io/gorm"
)

type FilmRepository struct {
	DB *gorm.DB
}

func (repo *FilmRepository) GetAllFilms() ([]dto.FilmDTO, error) {
	// Implementação de consulta ao banco
	return []dto.FilmDTO{}, nil
}
