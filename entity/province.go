package entity

import "gorm.io/gorm"

//Provinces Entity
type Provinces struct {
	gorm.Model
	Name     string      `json:"name"`
	District []Districts `gorm:"ForeignKey:ProvinceID"`
}
