package repository

import "github.com/SerafimKuzmin/sd/src/models"

type RepositoryI interface {
	CreateCookie(cookie *models.Cookie) error
	GetUserByCookie(value string) (string, error)
	DeleteCookie(value string) error
}
