package dto

import (
	"github.com/SerafimKuzmin/sd/backend/models"
	"time"
)

type ReqCreateUpdateList struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name" validate:"required"`
	CreateDT time.Time `json:"create_dt"`
}

func (req *ReqCreateUpdateList) ToModelList() *models.List {
	return &models.List{
		ID:       req.ID,
		Name:     req.Name,
		CreateDT: req.CreateDT,
	}
}

type RespList struct {
	ID       uint64    `json:"id"`
	Name     string    `json:"name"`
	CreateDT time.Time `json:"create_dt"`
}

func GetResponseFromModelList(list *models.List) *RespList {
	return &RespList{
		ID:       list.ID,
		Name:     list.Name,
		CreateDT: list.CreateDT,
	}
}

func GetResponseFromModelLists(lists []*models.List) []*RespList {
	result := make([]*RespList, 0, 10)
	for _, list := range lists {
		result = append(result, GetResponseFromModelList(list))
	}

	return result
}
