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
		SeedWareHouse(&db)
		SeedRack(&db)
		SeedCategory(&db)
		SeedUser(&db)
		SeedSupplier(&db)
	}
}
