package dto

import (
	"github.com/SerafimKuzmin/sd/backend/models"
	"time"
)

type ReqUserSignIn struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type ReqUserSignUp struct {
	Login     string    `json:"login" validate:"required"`
	Password  string    `json:"password" validate:"required"`
	Role      int       `json:"role" validate:"required"`
	Email     string    `json:"email" validate:"required"`
	FullName  string    `json:"full_name"`
	IsActive  bool      `json:"is_active"`
	CreateDT  time.Time `json:"create_dt"`
	CountryID *uint64   `json:"country_id"`
}

func (req *ReqUserSignIn) ToModelUser() *models.User {
	return &models.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *ReqUserSignUp) ToModelUser() *models.User {

	return &models.User{
		Login:     req.Login,
		Password:  req.Password,
		Role:      req.Role,
		Email:     req.Email,
		FullName:  req.FullName,
		IsActive:  req.IsActive,
		CreateDT:  req.CreateDT,
		CountryID: req.CountryID,
	}
}

type RespUser struct {
	ID        uint64    `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Role      int       `json:"role_id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	IsActive  bool      `json:"is_active"`
	CreateDT  time.Time `json:"create_dt"`
	CountryID *uint64   `json:"country_id"`
}

func GetResponseFromModelUser(user *models.User) *RespUser {
	return &RespUser{
		ID:        user.ID,
		Login:     user.Login,
		Password:  user.Password,
		Role:      user.Role,
		Email:     user.Email,
		FullName:  user.FullName,
		IsActive:  user.IsActive,
		CreateDT:  user.CreateDT,
		CountryID: user.CountryID,
	}
}
