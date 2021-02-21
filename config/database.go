package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var connDB *gorm.DB

//Connect ...
func Connect() {
	// var userDatabase, passDatabase, portDatabase, hostDatabase, nameDatabase string

	// userDatabase = os.Getenv("DB_USER")
	// passDatabase = os.Getenv("DB_PASSWORD")
	// portDatabase = os.Getenv("DB_PORT")
	// hostDatabase = os.Getenv("DB_HOST")
	// nameDatabase = os.Getenv("DB_NAME")

	// dsn :=
	// 	" host=" + hostDatabase +
	// 		" user=" + userDatabase +
	// 		" password=" + passDatabase +
	// 		" dbname=" + nameDatabase +
	// 		" port=" + portDatabase +
	// 		" sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open("postgres://iolrxtfzmvwbes:873fd93effab77843a7a71e5e8e98830e5f7222b2245aef2a314ac97f80cc516@ec2-34-198-31-223.compute-1.amazonaws.com:5432/dbhrucs3c9217j"), &gorm.Config{})
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
