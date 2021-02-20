package seeder

import (
	"fmt"
	"strconv"
	"warehouse/entity"

	"gorm.io/gorm"
)

// SeedTransaction func
func SeedTransaction(db *gorm.DB) {
	var transactionArray = [...][4]string{
		// product id, user id, rack id, stock
		{"1", "1", "1", "100"},
		{"3", "1", "7", "50"},
		{"2", "1", "6", "70"},
		{"4", "1", "11", "75"},
		{"5", "1", "1", "70"},
	}

	var transaction entity.Transactions
	for _, v := range transactionArray {
		product, _ := strconv.ParseUint(v[0], 10, 32)
		user, _ := strconv.ParseUint(v[1], 10, 32)
		rack, _ := strconv.ParseUint(v[2], 10, 32)
		stock, _ := strconv.ParseInt(v[3], 10, 32)

		transaction.ProductID = uint(product)
		transaction.UserID = uint(user)
		transaction.RackID = uint(rack)
		transaction.ProductStock = int(stock)

		transaction.ID = 0

		db.Create(&transaction)

	}
	fmt.Println("Seeder Transaction created Sucessfully")
}
