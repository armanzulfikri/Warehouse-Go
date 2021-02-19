package migrations

import (
	"os"
	"warehouse/config"
)

//Migrations ...
func Migrations()  {
	db := *config.GetConnection()
	if os.Getenv("USE_MIGRATE") == "yes" {
		provinceMigrations(&db)
	}
}
