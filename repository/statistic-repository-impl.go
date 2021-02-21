package repository

import (
	"warehouse/entity"
	"warehouse/model/response"

	"gorm.io/gorm"
)

//NewStatisticRepository ...
func NewStatisticRepository(database *gorm.DB) StatisticRepository {
	return &statisticRepositoryImpl{
		Database: database,
	}
}

type statisticRepositoryImpl struct {
	Database *gorm.DB
}

func (repository statisticRepositoryImpl) GetAll() (response response.StatisticResponse) {
	repository.Database.Model(&entity.Warehouses{}).Select("count(id) as warehouse").Scan(&response.Warehouse)
	repository.Database.Model(&entity.Products{}).Select("count(id) as product").Scan(&response.Product)
	repository.Database.Model(&entity.TransactionDetails{}).Select("count(id) as transaction").Scan(&response.Transaction)
	repository.Database.Model(&entity.Suppliers{}).Select("count(id) as supplier").Scan(&response.Supplier)

	return
}
