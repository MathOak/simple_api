package services

import (
	"simple_api/pkg/models/films/dto"
	"simple_api/pkg/models/films/repositories"
	"simple_api/pkg/models/films/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// O pacote services contém a lógica de negócio relacionada aos filmes.

// Importa os pacotes necessários:
// - dto: define os Data Transfer Objects para filmes.
// - repositories: fornece acesso ao repositório de filmes.
type FilmService struct {
	Repo *repositories.FilmRepository
}

func (s *FilmService) ListFilms() ([]dto.FilmDTO, error) {
	return s.Repo.GetAllFilms()
}

func (s *FilmService) ListAllFilms(c *gin.Context) ([]dto.FilmDTO, error) {
	// Recupera o *gorm.DB do contexto global (defina via middleware se necessário)
	db, ok := c.MustGet("db").(*gorm.DB)
	if !ok {
		c.JSON(500, gin.H{"error": "DB connection not found"})
		return
	}
	repo := &repositories.FilmRepository{DB: db}
	service := &services.FilmService{Repo: repo}
	films, err := service.ListFilms()

	return films, err
}
