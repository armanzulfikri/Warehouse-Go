package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//districtMigrations func
func districtMigrations(db *gorm.DB) {
	var checkTableDistricts bool

	checkTableDistricts = db.Migrator().HasTable(&entity.Districts{})
	if !checkTableDistricts {
		db.Migrator().CreateTable(&entity.Districts{})
		fmt.Println("Create Table Districts")
	}

}
