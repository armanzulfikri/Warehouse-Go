package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//categoryMigrations func
func categoryMigrations(db *gorm.DB) {
	var checkTableCategory bool

	checkTableCategory = db.Migrator().HasTable(&entity.Categories{})
	if !checkTableCategory {
		db.Migrator().CreateTable(&entity.Categories{})
		fmt.Println("Create Table CategorycheckTableCategory")
	}

}
