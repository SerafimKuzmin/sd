package dto

import (
	"github.com/SerafimKuzmin/sd/backend/models"
)

type ReqCreateUpdatePersonalRating struct {
	ID     uint64  `json:"id"`
	UserID uint64  `json:"user_id" validate:"required"`
	FilmID uint64  `json:"film_id" validate:"required"`
	Rate   float64 `json:"rating" validate:"required"`
}

func (req *ReqCreateUpdatePersonalRating) ToModelPersonalRating() *models.PersonalRating {
	return &models.PersonalRating{
		ID:     req.ID,
		UserID: req.UserID,
		FilmID: req.FilmID,
		Rate:   req.Rate,
	}
}

type RespPersonalRating struct {
	ID     uint64  `json:"id"`
	UserID uint64  `json:"user_id"`
	FilmID uint64  `json:"film_id"`
	Rate   float64 `json:"rating"`
}

func GetResponseFromModelPersonalRating(personalRating *models.PersonalRating) *RespPersonalRating {
	return &RespPersonalRating{
		ID:     personalRating.ID,
		UserID: personalRating.UserID,
		FilmID: personalRating.FilmID,
		Rate:   personalRating.Rate,
	}
}

func GetResponseFromModelPersonalRatings(PersonalRatings []*models.PersonalRating) []*RespPersonalRating {
	result := make([]*RespPersonalRating, 0, 10)
	for _, PersonalRating := range PersonalRatings {
		result = append(result, GetResponseFromModelPersonalRating(PersonalRating))
	}

	return result
}
