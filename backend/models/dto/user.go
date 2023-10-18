package dto

import (
	"github.com/SerafimKuzmin/sd/backend/models"
	"time"
)

type ReqUpdateUser struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Role      int       `json:"role_id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	IsActive  bool      `json:"is_active"`
	CreateDT  time.Time `json:"create_dt"`
	CountryID *uint64   `json:"country_id"`
}

func (req *ReqUpdateUser) ToModelUser() *models.User {
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

func GetResponseFromModelUsers(users []*models.User) []*RespUser {
	result := make([]*RespUser, 0, 10)
	for _, user := range users {
		result = append(result, GetResponseFromModelUser(user))
	}

	return result
}
