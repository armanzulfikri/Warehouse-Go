package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//productMigrations func
func productMigrations(db *gorm.DB) {
	var checkTableProduct bool

	db.Migrator().DropTable(&entity.Products{})

	checkTableProduct = db.Migrator().HasTable(&entity.Products{})
	if !checkTableProduct {
		db.Migrator().CreateTable(&entity.Products{})
		fmt.Println("Create Table product")
	}
}
