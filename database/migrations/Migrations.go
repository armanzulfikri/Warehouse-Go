package migrations

import (
	"os"
	"warehouse/config"
)

//Migrations ...
func Migrations() {
	db := *config.GetConnection()
	if os.Getenv("USE_MIGRATE") == "yes" {
		provinceMigrations(&db)
		districtMigrations(&db)
		rackMigrations(&db)
		warehousesMigrations(&db)
		categoryMigrations(&db)
		supplierMigrations(&db)
		userMigrations(&db)
	}
}
