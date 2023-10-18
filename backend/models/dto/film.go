package dto

import (
	"github.com/SerafimKuzmin/sd/backend/models"
	"time"
)

type ReqCreateUpdateFilm struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Rate        float64   `json:"rate"`
	ReleaseDT   time.Time `json:"release_dt"`
	Duration    uint      `json:"duration"`
}

func (req *ReqCreateUpdateFilm) ToModelFilm() *models.Film {
	return &models.Film{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Rate:        req.Rate,
		ReleaseDT:   req.ReleaseDT,
		Duration:    req.Duration,
	}
}

type RespFilm struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rate        float64   `json:"rate"`
	ReleaseDT   time.Time `json:"release_dt"`
	Duration    uint      `json:"duration"`
}

func GetResponseFromModelFilm(film *models.Film) *RespFilm {
	return &RespFilm{
		ID:          film.ID,
		Name:        film.Name,
		Description: film.Description,
		Rate:        film.Rate,
		ReleaseDT:   film.ReleaseDT,
		Duration:    film.Duration,
	}
}

func GetResponseFromModelFilms(Films []*models.Film) []*RespFilm {
	result := make([]*RespFilm, 0, 10)
	for _, Film := range Films {
		result = append(result, GetResponseFromModelFilm(Film))
	}

	return result
}
