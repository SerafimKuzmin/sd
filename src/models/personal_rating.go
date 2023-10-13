package models

type PersonalRating struct {
	ID     uint64  `json:"id"`
	UserID uint64  `json:"user_id"`
	FilmID uint64  `json:"film_id"`
	Rate   float64 `json:"rate"`
}
