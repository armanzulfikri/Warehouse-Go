package migrations

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

//transactionDetailMigrations func
func transactionDetailMigrations(db *gorm.DB) {
	var checkTableTransactionDetail bool

	db.Migrator().DropTable(&entity.TransactionDetails{})

	checkTableTransactionDetail = db.Migrator().HasTable(&entity.TransactionDetails{})
	if !checkTableTransactionDetail {
		db.Migrator().CreateTable(&entity.TransactionDetails{})
		fmt.Println("Create Table TransactionDetails")
	}
}
