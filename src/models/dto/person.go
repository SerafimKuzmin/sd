package dto

import (
	"github.com/SerafimKuzmin/sd/src/models"
)

type ReqCreateUpdatePerson struct {
	ID   uint64 `json:"id"`
	Name string `json:"name" validate:"required"`
}

func (req *ReqCreateUpdatePerson) ToModelPerson() *models.Person {
	return &models.Person{
		ID:   req.ID,
		Name: req.Name,
	}
}

type RespPerson struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

func GetResponseFromModelPerson(person *models.Person) *RespPerson {
	return &RespPerson{
		ID:   person.ID,
		Name: person.Name,
	}
}

func GetResponseFromModelPersons(persons []*models.Person) []*RespPerson {
	result := make([]*RespPerson, 0, 10)
	for _, person := range persons {
		result = append(result, GetResponseFromModelPerson(person))
	}

	return result
}
