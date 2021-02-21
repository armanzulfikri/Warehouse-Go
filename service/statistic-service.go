package service

import "warehouse/model/response"

//StatisticService ...
type StatisticService interface {
	List() (responses response.StatisticResponse)
}
