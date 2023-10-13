package models

import (
	"time"
)

type PersonalRating struct {
	ID        uint64    `json:"id"`
	UserID    uint64    `json:"user_id"`
	FilmID    uint64    `json:"film_id"`
	Rating    string    `json:"rating"`
	RatingDT  time.Time `json:"rating_dt"`
	Review    string    `json:"review"`
	IsDeleted string    `json:"is_deleted"`
}
