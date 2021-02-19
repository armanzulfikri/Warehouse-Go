package seeder

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

// SeedRack func
func SeedRack(db *gorm.DB) {
	var categoryArray = [...][4]string{
		{"1", "1", "KLTNF-01", "100"},
		{"2", "3", "JMBRM-01", "300"},
		{"3", "5", "GRSKI-01", "100"},
		{"4", "21", "BNDGF-01", "400"},
		{"5", "13", "BNDGK-01", "300"},
		{"6", "2", "LMPNGF-01", "100"},
		{"7", "3", "PDNGM-01", "200"},
		{"8", "6", "MJKRTI-01", "100"},
		{"9", "20", "MNHSK-01", "250"},
		{"10", "11", "MDNF-01", "80"},
		{"11", "2", "MLNGF-01", "80"},
	}

	var category entity.Categories
	for _, v := range categoryArray {
		category.CategoryName = v[0]

		category.ID = 0

		db.Create(&category)

	}
	fmt.Println("Seeder Rack created Sucessfully")
}
