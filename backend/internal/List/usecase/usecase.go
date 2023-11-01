package usecase

import (
	"github.com/SerafimKuzmin/sd/backend/models"
	"github.com/pkg/errors"
)

type UsecaseI interface {
	CreateList(t *models.List) error
	DeleteList(listId uint64) error
	UpdateList(t *models.List) error
	GetList(listId uint64) (*models.List, error)
	AddFilm(listID uint64, filmID uint64) error
	GetUserLists(userID uint64) ([]*models.List, error)
	GetFilmsByList(listID uint64) ([]*models.Film, error)
}

type usecase struct {
	listRepository RepositoryI
}

func New(lRep RepositoryI) UsecaseI {
	return &usecase{
		listRepository: lRep,
	}
}

func (u *usecase) CreateList(t *models.List) error {
	err := u.listRepository.CreateList(t)

	if err != nil {
		return errors.Wrap(err, "Error in func List.Usecase.CreateList")
	}

	return nil
}

func (u *usecase) GetList(id uint64) (*models.List, error) {
	resTag, err := u.listRepository.GetList(id)

	if err != nil {
		return nil, errors.Wrap(err, "Tag.usecase.GetTag error while get Tag info")
	}

	return resTag, nil
}

func (u *usecase) UpdateList(t *models.List) error {
	err := u.listRepository.UpdateList(t)

	existedTag, err := u.listRepository.GetList(t.ID)
	if err != nil {
		return err
	}

	if existedTag == nil {
		return errors.New("List not found") //TODO models error
	}

	if err != nil {
		return errors.Wrap(err, "Error in func List.Usecase.CreateList")
	}

	return nil
}

func (u *usecase) DeleteList(listId uint64) error {
	existedTag, err := u.listRepository.GetList(listId)
	if err != nil {
		return err
	}

	if existedTag == nil {
		return errors.New("List not found") //TODO models error
	}

	err = u.listRepository.DeleteList(listId)

	if err != nil {
		return errors.Wrap(err, "Tag.repository delete error")
	}

	return nil
}

func (u *usecase) AddFilm(listID uint64, filmID uint64) error {
	err := u.listRepository.AddFilm(listID, filmID)

	if err != nil {
		return errors.Wrap(err, "Error in func List.Usecase.CreateList")
	}

	return nil
}

func (u *usecase) GetUserLists(userID uint64) ([]*models.List, error) {
	entries, err := u.listRepository.GetUserLists(userID)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func Tag.Usecase.GetUserPosts")
	}

	return entries, nil
}

func (u *usecase) GetFilmsByList(listID uint64) ([]*models.Film, error) {
	entries, err := u.listRepository.GetFilmsByList(listID)

	if err != nil {
		return nil, errors.Wrap(err, "Error in func Tag.Usecase.GetUserPosts")
	}

	return entries, nil
}
