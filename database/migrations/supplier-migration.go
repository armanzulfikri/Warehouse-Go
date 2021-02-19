package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//supplierMigrations func
func supplierMigrations(db *gorm.DB) {
	var checkTableSuppliers bool

	checkTableSuppliers = db.Migrator().HasTable(&entity.Suppliers{})
	if !checkTableSuppliers {
		db.Migrator().CreateTable(&entity.Suppliers{})
		fmt.Println("Create Table Suppliers")
	}

}
