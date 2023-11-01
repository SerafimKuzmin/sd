package usecase

import (
	"github.com/SerafimKuzmin/sd/backend/models"
)

type RepositoryI interface {
	CreatePersonalRating(e *models.PersonalRating) error
	UpdatePersonalRating(e *models.PersonalRating) error
	GetPersonalRating(id uint64) (*models.PersonalRating, error)
	DeletePersonalRating(id uint64) error
}
