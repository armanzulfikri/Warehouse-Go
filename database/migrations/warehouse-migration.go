package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//warehousesMigrations func
func warehousesMigrations(db *gorm.DB) {
	var checkTableWarehouses bool

	checkTableWarehouses = db.Migrator().HasTable(&entity.Warehouses{})
	if !checkTableWarehouses {
		db.Migrator().CreateTable(&entity.Warehouses{})
		fmt.Println("Create Table Warehouses")
	}

}
