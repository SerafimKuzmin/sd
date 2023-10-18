package models

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Role      int       `json:"role_id"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
	IsActive  bool      `json:"is_active"`
	CreateDT  time.Time `json:"create_dt"`
	CountryID uint64    `json:"country_id"`
}
