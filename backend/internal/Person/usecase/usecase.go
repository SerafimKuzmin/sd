package usecase

import (
	"github.com/SerafimKuzmin/sd/backend/internal/cache"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
)

type UsecaseI interface {
	CreatePerson(e *models.Person) error
	UpdatePerson(e *models.Person) error
	GetPerson(id uint64) (*models.Person, error)
	DeletePerson(id uint64) error
}

type usecase struct {
	PersonRepository RepositoryI
	redisStorage     cache.CacheStorageI
}

func New(pRep RepositoryI, rS cache.CacheStorageI) UsecaseI {
	return &usecase{
		PersonRepository: pRep,
		redisStorage:     rS,
	}
}

func (u *usecase) CreatePerson(e *models.Person) error {
	err := u.PersonRepository.CreatePerson(e)

	if err != nil {
		return errors.Wrap(err, "Error in func Person.Usecase.CreatePerson")
	}

	return nil
}

func (u *usecase) UpdatePerson(p *models.Person) error {
	_, err := u.PersonRepository.GetPerson(p.ID)

	if err != nil {
		return errors.Wrap(err, "Error in func Person.Usecase.Update.GetPerson")
	}

	err = u.PersonRepository.UpdatePerson(p)

	if err != nil {
		return errors.Wrap(err, "Error in func Person.Usecase.Update")
	}

	return nil
}

func (u *usecase) GetPerson(id uint64) (*models.Person, error) {
	resPerson, err := u.PersonRepository.GetPerson(id)

	if err != nil {
		return nil, errors.Wrap(err, "Person.usecase.GetPerson error while get Person info")
	}

	return resPerson, nil
}

func (u *usecase) DeletePerson(id uint64) error {
	existedPerson, err := u.PersonRepository.GetPerson(id)
	if err != nil {
		return err
	}

	if existedPerson == nil {
		return errors.New("Person not found") //TODO models error
	}

	//if *existedPerson.UserID != userID {
	//	return errors.New("Permission denied")
	//}

	err = u.PersonRepository.DeletePerson(id)

	if err != nil {
		return errors.Wrap(err, "Error in func Person.Usecase.DeletePerson repository")
	}

	return nil
}
