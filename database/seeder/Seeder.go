package seeder

import (
	"os"
	"warehouse/config"
)

//Seeder ...
func Seeder() {
	db := *config.GetConnection()
	if os.Getenv("USE_SEEDER") == "yes" {
		SeedProvince(&db)
		SeedDisctrict(&db)
		SeedUser(&db)
		SeedCategory(&db)
		SeedWareHouse(&db)
		SeedRack(&db)
	}
}
