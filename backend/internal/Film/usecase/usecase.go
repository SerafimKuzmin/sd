package usecase

import (
	FilmRep "github.com/SerafimKuzmin/sd/backend/internal/Film/repository"
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
)

type UsecaseI interface {
	CreateFilm(g *models.Film) error
	UpdateFilm(g *models.Film) error
	GetFilm(id uint64) (*models.Film, error)
	DeleteFilm(id uint64) error
	GetFilmByPerson(id uint64) ([]*models.Film, error)
	GetFilmByCountry(id uint64) ([]*models.Film, error)
}

type usecase struct {
	filmRepository FilmRep.RepositoryI
}

func New(gRep FilmRep.RepositoryI) UsecaseI {
	return &usecase{
		filmRepository: gRep,
	}
}

func (u *usecase) CreateFilm(e *models.Film) error {
	err := u.filmRepository.CreateFilm(e)

	if err != nil {
		return errors.Wrap(err, "Error in func Film.Usecase.CreateFilm")
	}

	return nil
}

func (u *usecase) UpdateFilm(Film *models.Film) error {
	_, err := u.filmRepository.GetFilm(Film.ID)

	if err != nil {
		return errors.Wrap(err, "Error in func Film.Usecase.Update.GetFilm")
	}

	err = u.filmRepository.UpdateFilm(Film)

	if err != nil {
		return errors.Wrap(err, "Error in func Film.Usecase.CreateFilm")
	}

	return nil
}

func (u *usecase) GetFilm(id uint64) (*models.Film, error) {
	resFilm, err := u.filmRepository.GetFilm(id)

	if err != nil {
		return nil, errors.Wrap(err, "Film.usecase.GetFilm error while get Film info")
	}

	return resFilm, nil
}

func (u *usecase) DeleteFilm(id uint64) error {
	existedFilm, err := u.filmRepository.GetFilm(id)
	if err != nil {
		return err
	}

	if existedFilm == nil {
		return errors.New("Film not found") //TODO models error
	}

	err = u.filmRepository.DeleteFilm(id)

	if err != nil {
		return errors.Wrap(err, "Error in func Film.Usecase.DeleteFilm repository")
	}

	return nil
}

func (u *usecase) GetFilmByPerson(id uint64) ([]*models.Film, error) {
	entries, err := u.filmRepository.GetFilmByPerson(id)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func Tag.Usecase.GetUserPosts")
	}

	return entries, nil
}

func (u *usecase) GetFilmByCountry(id uint64) ([]*models.Film, error) {
	entries, err := u.filmRepository.GetFilmByCountry(id)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func Tag.Usecase.GetUserPosts")
	}

	return entries, nil
}
