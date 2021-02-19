package entity

import (
	"time"

	"gorm.io/gorm"
)

//Provinces Entity
type Provinces struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `json:"name"`
	District  []Districts    `gorm:"ForeignKey:ProvinceID"`
}
