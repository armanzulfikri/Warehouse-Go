package config

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"os"
)

var connDB *gorm.DB

//Connect ...
func Connect() {
	var userDatabase, passDatabase, portDatabase, hostDatabase, nameDatabase string

	userDatabase = os.Getenv("DB_USER")
	passDatabase = os.Getenv("DB_PASSWORD")
	portDatabase = os.Getenv("DB_PORT")
	hostDatabase = os.Getenv("DB_HOST")
	nameDatabase = os.Getenv("DB_NAME")

	dsn :=
		" host=" + hostDatabase +
			" user=" + userDatabase +
			" password=" + passDatabase +
			" dbname=" + nameDatabase +
			" port=" + portDatabase +
			" sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	} else {
		fmt.Println("Koneksi Sukses")
	}

	connDB = db
}

//GetConnection ...
func GetConnection() *gorm.DB {
	return connDB
}