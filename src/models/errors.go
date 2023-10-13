package models

import "errors"

var (
	ErrNotFound            = errors.New("item is not found")
	ErrInvalidPassword     = errors.New("invalid password")
	ErrConflictNickname    = errors.New("nickname already exists")
	ErrConflictEmail       = errors.New("email already exists")
	ErrBadRequest          = errors.New("bad request")
	ErrConflictFilm        = errors.New("film already exists")
	ErrUnauthorized        = errors.New("no cookie")
	ErrInternalServerError = errors.New("internal server error")
	ErrPermissionDenied    = errors.New("permission denied")
)
