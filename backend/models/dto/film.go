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
	Genre       string    `json:"genre"`
	ReleaseDT   time.Time `json:"release_dt"`
	Duration    uint      `json:"duration"`
	CountryID   *uint64   `json:"country_id"`
}

func (req *ReqCreateUpdateFilm) ToModelFilm() *models.Film {
	return &models.Film{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Rate:        req.Rate,
		Genre:       req.Genre,
		ReleaseDT:   req.ReleaseDT,
		Duration:    req.Duration,
		CountryID:   req.CountryID,
	}
}

type RespFilm struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Rate        float64   `json:"rate"`
	Genre       string    `json:"genre"`
	ReleaseDT   time.Time `json:"release_dt"`
	Duration    uint      `json:"duration"`
	CountryID   *uint64   `json:"country_id"`
}

func GetResponseFromModelFilm(film *models.Film) *RespFilm {
	return &RespFilm{
		ID:          film.ID,
		Name:        film.Name,
		Description: film.Description,
		Rate:        film.Rate,
		Genre:       film.Genre,
		ReleaseDT:   film.ReleaseDT,
		Duration:    film.Duration,
		CountryID:   film.CountryID,
	}
}

func GetResponseFromModelFilms(Films []*models.Film) []*RespFilm {
	result := make([]*RespFilm, 0, 10)
	for _, Film := range Films {
		result = append(result, GetResponseFromModelFilm(Film))
	}

	return result
}
