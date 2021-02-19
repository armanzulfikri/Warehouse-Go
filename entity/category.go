package entity

import "gorm.io/gorm"

//Categories Entity
type Categories struct {
	gorm.Model
	CategoryName string `json:"category_name"`
}
