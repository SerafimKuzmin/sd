package usecase

import "github.com/SerafimKuzmin/sd/backend/models"

type RepositoryI interface {
	CreateCookie(cookie *models.Cookie) error
	GetUserByCookie(value string) (string, error)
	DeleteCookie(value string) error
}
