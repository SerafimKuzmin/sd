package models

import (
	"time"
)

type RoleType int64

const (
	Admin RoleType = iota
	Moderator
	DefaultUser
)

func (s RoleType) String() string {
	switch s {
	case DefaultUser:
		return "user"
	case Admin:
		return "admin"
	}
	return "unknown"
}

type User struct {
	ID        uint64    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"username"`
	Password  string    `json:"password"`
	BirthDate time.Time `json:"birth_date"`
	Sex       bool      `json:"sex"`
	CountryID uint64    `json:"country_id"`
	IsDeleted string    `json:"is_deleted"`
	IsActive  string    `json:"is_active"`
}
