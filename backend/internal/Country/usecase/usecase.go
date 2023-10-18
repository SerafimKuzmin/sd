package usecase

import (
	CountryRep "github.com/SerafimKuzmin/sd/backend/internal/Country/repository"
	"github.com/SerafimKuzmin/sd/backend/internal/cache"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
)

type UsecaseI interface {
	CreateCountry(e *models.Country) error
	UpdateCountry(e *models.Country) error
	GetCountry(id uint64) (*models.Country, error)
	DeleteCountry(id uint64) error
}

type usecase struct {
	CountryRepository CountryRep.RepositoryI
	redisStorage      cache.CacheStorageI
}

func New(pRep CountryRep.RepositoryI, rS cache.CacheStorageI) UsecaseI {
	return &usecase{
		CountryRepository: pRep,
		redisStorage:      rS,
	}
}

func (u *usecase) CreateCountry(e *models.Country) error {
	err := u.CountryRepository.CreateCountry(e)

	if err != nil {
		return errors.Wrap(err, "Error in func Country.Usecase.CreateCountry")
	}

	return nil
}

func (u *usecase) UpdateCountry(p *models.Country) error {
	_, err := u.CountryRepository.GetCountry(p.ID)

	if err != nil {
		return errors.Wrap(err, "Error in func Country.Usecase.Update.GetCountry")
	}

	err = u.CountryRepository.UpdateCountry(p)

	if err != nil {
		return errors.Wrap(err, "Error in func Country.Usecase.Update")
	}

	return nil
}

func (u *usecase) GetCountry(id uint64) (*models.Country, error) {
	resCountry, err := u.CountryRepository.GetCountry(id)

	if err != nil {
		return nil, errors.Wrap(err, "Country.usecase.GetCountry error while get Country info")
	}

	return resCountry, nil
}

func (u *usecase) DeleteCountry(id uint64) error {
	existedCountry, err := u.CountryRepository.GetCountry(id)
	if err != nil {
		return err
	}

	if existedCountry == nil {
		return errors.New("Country not found") //TODO models error
	}

	//if *existedCountry.UserID != userID {
	//	return errors.New("Permission denied")
	//}

	err = u.CountryRepository.DeleteCountry(id)

	if err != nil {
		return errors.Wrap(err, "Error in func Country.Usecase.DeleteCountry repository")
	}

	return nil
}
