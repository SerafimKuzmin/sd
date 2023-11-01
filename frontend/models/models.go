package models

import "time"

type ErrorBody struct {
	Err string `json:"error"`
}

type FilmList struct {
	FilmID uint64
	ListID uint64
}

type RateFilm struct {
	ID     uint64
	UserID uint64
	FilmID uint64
	Rate   float64
}

type Film struct {
	ID          uint64
	Name        string
	Description string
	Rate        float64
	Genre       string
	ReleaseDT   time.Time
	Duration    uint
}

type User struct {
	ID        uint64
	Login     string
	Password  string
	Role      int
	Email     string
	FullName  string
	IsActive  bool
	CreateDT  time.Time
	CountryID uint64
}

type List struct {
	ID       uint64
	Name     string
	CreateDT time.Time
}

type Lists struct {
	Lists []struct {
		ID       uint64    `json:"id"`
		Name     string    `json:"name"`
		CreateDT time.Time `json:"create_dt"`
	} `json:"body"`
}

type FilmsResponse struct {
	Films []struct {
		ID          uint64    `json:"id"`
		Name        string    `json:"name"`
		Description string    `json:"description"`
		Rate        float64   `json:"rate"`
		ReleaseDT   time.Time `json:"release_dt"`
		Duration    uint      `json:"duration"`
	} `json:"body"`
}

type FilmResponse struct {
	Body Film
}

type Client struct {
	ClientId uint64 `json:"ClientId"`
	Login    string `json:"Login"`
	Password string
	Token    string `json:"Token"`
}
