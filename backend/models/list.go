package models

import (
	"time"
)

type List struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	CreateDT time.Time `json:"create_dt"`
}

type ListFilm struct {
	ID     uint64 `json:"id"`
	FilmID uint64 `json:"film_id"`
}
