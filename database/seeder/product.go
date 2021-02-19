package seeder

import (
	"fmt"
	"strconv"
	"warehouse/entity"

	"gorm.io/gorm"
)

// SeedProduct func
func SeedProduct(db *gorm.DB) {
	var productArray = [...][5]string{
		{"1", "4", "2", "Indomie Goreng", "https://images.unsplash.com/photo-1612929633738-8fe44f7ec841?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=634&q=80"},
		{"1", "5", "1", "Biskuat", "https://images.unsplash.com/photo-1608769240116-a45a6ac33a55?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1842&q=80"},
		{"1", "3", "3", "Energen", "https://images.unsplash.com/photo-1559598467-f8b76c8155d0?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1351&q=80"},
		{"1", "4", "4", "Mie Sedaap", "https://images.unsplash.com/photo-1594100889583-a421e2b8265c?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1350&q=80"},
		{"1", "4", "5", "Gudang Garam Filter", "https://images.unsplash.com/photo-1527099908998-5b73a5fe2a0d?ixid=MXwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHw%3D&ixlib=rb-1.2.1&auto=format&fit=crop&w=1355&q=80"},
	}

	var product entity.Products
	for _, v := range productArray {
		user, _ := strconv.ParseUint(v[0], 10, 32)
		category, _ := strconv.Atoi(v[1])
		supplier, _ := strconv.Atoi(v[2])

		product.UserID = uint(user)
		product.CategoryID = uint(category)
		product.SupplierID = uint(supplier)
		product.ProductName = v[3]
		product.ProductImageURL = v[4]

		product.ID = 0

		db.Create(&product)

	}
	fmt.Println("Seeder Product created Sucessfully")
}
