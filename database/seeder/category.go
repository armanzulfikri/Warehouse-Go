package seeder

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

// SeedCategory func
func SeedCategory(db *gorm.DB) {
	var categoryArray = [...][1]string{
		{"Makanan Beku"},
		{"Makanan Kering"},
		{"Minuman Instan"},
		{"Makanan Instan"},
		{"Makanan Ringan"},
		{"Kayu"}, {"Besi"}, {"Kaca"}, {"Keramik"},
		{"Ram"}, {"Motherboard"}, {"Processor"}, {"Graphic Card"}, {"Handphone"},
		{"Jaket"}, {"Kemeja"}, {"Kaos"}, {"Celana"}, {"Kaos Kaki"}, {"Topi"},
		{"Masker"}, {"Serum"}, {"Eye Liner"}, {"Bedak"}, {"Lipstik"},
		{"Obat"}, {"Vitamin"}, {"Hand Sanitizer"}, {"Tisu"},
	}

	var category entity.Categories
	for _, v := range categoryArray {
		category.CategoryName = v[0]

		category.ID = 0

		db.Create(&category)

	}
	fmt.Println("Seeder Category created Sucessfully")
}
