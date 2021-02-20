package seeder

import (
	"fmt"
	"strconv"
	"warehouse/entity"

	"gorm.io/gorm"
)

// SeedRack func
func SeedRack(db *gorm.DB) {
	var SeedRackArray = [...][4]string{
		{"1", "1", "KLTNF-01", "100"},
		{"2", "3", "JMBRM-01", "300"},
		{"3", "5", "GRSKI-01", "100"},
		{"4", "21", "BNDGF-01", "400"},
		{"5", "13", "BNDGK-01", "300"},
		{"6", "2", "LMPNGF-01", "100"},
		{"7", "3", "PDNGM-01", "200"},
		{"8", "6", "MJKRTI-01", "100"},
		{"9", "20", "MNHSK-01", "250"},
		{"10", "11", "MDNF-01", "80"},
		{"11", "2", "MLNGF-01", "80"},
	}

	var racks entity.Racks
	for _, v := range SeedRackArray {
		warehouse, _ := strconv.ParseUint(v[0], 10, 32)
		category, _ := strconv.Atoi(v[1])
		capacity, _ := strconv.Atoi(v[3])

		racks.WarehouseID = uint(warehouse)
		racks.CategoryID = uint(category)
		racks.RackCode = v[2]
		racks.RackCapacity = int(capacity)

		racks.ID = 0

		db.Create(&racks)

	}
	fmt.Println("Seeder Rack created Sucessfully")
}
