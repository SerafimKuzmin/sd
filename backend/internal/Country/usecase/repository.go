package usecase

import (
	"github.com/SerafimKuzmin/sd/backend/models"
)

type RepositoryI interface {
	CreateCountry(e *models.Country) error
	UpdateCountry(e *models.Country) error
	GetCountry(id uint64) (*models.Country, error)
	DeleteCountry(id uint64) error
}
