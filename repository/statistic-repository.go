package repository

import "warehouse/model/response"

//StatisticRepository ...
type StatisticRepository interface {
	GetAll() (resp response.StatisticResponse)
}
