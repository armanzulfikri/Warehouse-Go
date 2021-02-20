package seeder

import (
	"fmt"
	"strconv"
	"warehouse/entity"

	"gorm.io/gorm"
)

// SeedTransactionDetail func
func SeedTransactionDetail(db *gorm.DB) {
	var transactionDetailArray = [...][5]string{
		// transaction id, product code, date of type, date, qty
		{"1", "INDF-01", "date_in", "100"},
		{"2", "ENRG-01", "date_in", "50"},
		{"3", "BSKT-01", "date_in", "70"},
		{"4", "MSDP-01", "date_in", "75"},
		{"5", "GGFT-01", "date_in", "70"},
	}

	var detail entity.TransactionDetails
	for _, v := range transactionDetailArray {
		transaction, _ := strconv.ParseUint(v[0], 10, 32)
		qty, _ := strconv.ParseInt(v[3], 10, 32)

		detail.TransactionID = uint(transaction)
		detail.ProductCode = v[1]
		detail.DateOfType = v[2]
		detail.Quantity = int(qty)

		detail.ID = 0

		db.Create(&detail)

	}
	fmt.Println("Seeder Transaction created Sucessfully")
}
