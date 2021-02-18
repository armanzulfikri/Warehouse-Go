package config

import (
	"gorm.io/gorm"
)

var connDB *gorm.DB

func Connect() {

}

func GetConnection() *gorm.DB {
	return connDB
}