package repository

import "github.com/SerafimKuzmin/sd/backend/models"

type RepositoryI interface {
	CreateList(t *models.List) error
	DeleteList(listId uint64) error
	UpdateList(t *models.List) error
	GetList(listId uint64) (*models.List, error)
	AddFilm(listID uint64, filmID uint64) error
	GetUserLists(userID uint64) ([]*models.List, error)
	GetFilmsByList(listID uint64) ([]*models.Film, error)
}
