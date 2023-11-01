package usecase

import (
	"github.com/SerafimKuzmin/sd/backend/models"
)

type RepositoryI interface {
	CreatePerson(e *models.Person) error
	UpdatePerson(e *models.Person) error
	GetPerson(id uint64) (*models.Person, error)
	DeletePerson(id uint64) error
}
