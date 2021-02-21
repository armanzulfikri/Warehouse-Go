package service

import (
	"warehouse/model/response"
	"warehouse/repository"
)

//NewStatisticService ...
func NewStatisticService(statisticRepository *repository.StatisticRepository) StatisticService {
	return &statisticServiceImpl{
		Repository: *statisticRepository,
	}
}

type statisticServiceImpl struct {
	Repository repository.StatisticRepository
}

//List
func (service statisticServiceImpl) List() (responses response.StatisticResponse) {
	responses = service.Repository.GetAll()

	return responses
}
