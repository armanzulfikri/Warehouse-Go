package entity

import "gorm.io/gorm"

//Suppliers struct
type Suppliers struct {
	gorm.Model
	SupplierName string `json:"supplier_name"`
	Address      string `json:"address"`
	Telepon      string `json:"role"`
}
