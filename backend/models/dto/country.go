package dto

import (
	"github.com/SerafimKuzmin/sd/backend/models"
)

type ReqCreateUpdateCountry struct {
	ID   uint64 `json:"id"`
	Name string `json:"name" validate:"required"`
}

func (req *ReqCreateUpdateCountry) ToModelCountry() *models.Country {
	return &models.Country{
		ID:   req.ID,
		Name: req.Name,
	}
}

type RespCountry struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func GetResponseFromModelCountry(country *models.Country) *RespCountry {
	return &RespCountry{
		ID:   country.ID,
		Name: country.Name,
	}
}

func GetResponseFromModelCountries(Countries []*models.Country) []*RespCountry {
	result := make([]*RespCountry, 0, 10)
	for _, Country := range Countries {
		result = append(result, GetResponseFromModelCountry(Country))
	}

	return result
}
