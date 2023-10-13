package repository

import (
	"github.com/SerafimKuzmin/sd/src/models"
)

type RepositoryI interface {
	CreateCountry(e *models.Country) error
	UpdateCountry(e *models.Country) error
	GetCountry(id uint64) (*models.Country, error)
	DeleteCountry(id uint64) error
}
