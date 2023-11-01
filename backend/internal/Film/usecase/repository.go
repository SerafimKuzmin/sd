package usecase

import (
	"github.com/SerafimKuzmin/sd/backend/models"
)

type RepositoryI interface {
	CreateFilm(g *models.Film) error
	UpdateFilm(g *models.Film) error
	GetFilm(id uint64) (*models.Film, error)
	DeleteFilm(id uint64) error
	GetFilmByPerson(id uint64) ([]*models.Film, error)
	GetFilmByCountry(id uint64) ([]*models.Film, error)
}
