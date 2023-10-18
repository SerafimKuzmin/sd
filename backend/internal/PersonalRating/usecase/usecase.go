package usecase

import (
	PersonalRatingRep "github.com/SerafimKuzmin/sd/backend/internal/PersonalRating/repository"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
)

type UsecaseI interface {
	CreatePersonalRating(t *models.PersonalRating) error
	UpdatePersonalRating(t *models.PersonalRating) error
	GetPersonalRating(id uint64) (*models.PersonalRating, error)
	DeletePersonalRating(id uint64, userID uint64) error
}

type usecase struct {
	PersonalRatingRepository PersonalRatingRep.RepositoryI
}

func (u *usecase) CreatePersonalRating(t *models.PersonalRating) error {
	err := u.PersonalRatingRepository.CreatePersonalRating(t)

	if err != nil {
		return errors.Wrap(err, "Error in func PersonalRating.Usecase.CreatePersonalRating")
	}

	return nil
}

func (u *usecase) UpdatePersonalRating(t *models.PersonalRating) error {
	_, err := u.PersonalRatingRepository.GetPersonalRating(t.ID)

	if err != nil {
		return errors.Wrap(err, "Error in func PersonalRating.Usecase.CreatePersonalRating")
	}

	err = u.PersonalRatingRepository.UpdatePersonalRating(t)

	if err != nil {
		return errors.Wrap(err, "Error in func PersonalRating.Usecase.CreatePersonalRating")
	}

	return nil
}

func (u *usecase) GetPersonalRating(id uint64) (*models.PersonalRating, error) {
	resPersonalRating, err := u.PersonalRatingRepository.GetPersonalRating(id)

	if err != nil {
		return nil, errors.Wrap(err, "PersonalRating.usecase.GetPersonalRating error while get PersonalRating info")
	}

	return resPersonalRating, nil
}

func (u *usecase) DeletePersonalRating(id uint64, userID uint64) error {
	existedPersonalRating, err := u.PersonalRatingRepository.GetPersonalRating(id)
	if err != nil {
		return err
	}

	if existedPersonalRating == nil {
		return errors.New("PersonalRating not found") //TODO models error
	}

	if existedPersonalRating.UserID != userID {
		return errors.New("Permission denied")
	}

	err = u.PersonalRatingRepository.DeletePersonalRating(id)

	if err != nil {
		return errors.Wrap(err, "PersonalRating.repository delete error")
	}

	return nil
}

func New(tRep PersonalRatingRep.RepositoryI) UsecaseI {
	return &usecase{
		PersonalRatingRepository: tRep,
	}
}
