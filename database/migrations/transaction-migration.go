package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//transactionMigrations func
func transactionMigrations(db *gorm.DB) {
	var checkTableTransaction bool

	db.Migrator().DropTable(&entity.Transactions{})

	checkTableTransaction = db.Migrator().HasTable(&entity.Transactions{})
	if !checkTableTransaction {
		db.Migrator().CreateTable(&entity.Transactions{})
		fmt.Println("Create Table transaction")
	}
}
