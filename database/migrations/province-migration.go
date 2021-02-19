package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

func provinceMigrations(db *gorm.DB) {
	var checkTableProvinces bool

	checkTableProvinces = db.Migrator().HasTable(&entity.Provinces{})
	if !checkTableProvinces {
		db.Migrator().CreateTable(&entity.Provinces{})
		fmt.Println("Create Table Provinces")
	}

}
