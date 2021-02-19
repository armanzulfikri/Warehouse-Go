package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//rackMigrations func
func rackMigrations(db *gorm.DB) {
	var checkTableRacks bool

	db.Migrator().DropTable(&entity.Racks{})

	checkTableRacks = db.Migrator().HasTable(&entity.Racks{})
	if !checkTableRacks {
		db.Migrator().CreateTable(&entity.Racks{})
		fmt.Println("Create Table Racks")
	}

}
