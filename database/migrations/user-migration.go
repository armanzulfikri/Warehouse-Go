package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//userMigrations func
func userMigrations(db *gorm.DB) {
	var checkTableUsers bool

	db.Migrator().DropTable(&entity.Users{})

	checkTableUsers = db.Migrator().HasTable(&entity.Users{})
	if !checkTableUsers {
		db.Migrator().CreateTable(&entity.Users{})
		fmt.Println("Create Table Users")
	}

}
