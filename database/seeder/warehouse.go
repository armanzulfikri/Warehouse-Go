package seeder

import (
	"fmt"
	"strconv"
	"warehouse/entity"

	"gorm.io/gorm"
)

//SeedWareHouse func
func SeedWareHouse(db *gorm.DB) {
	var SeedWareHouseArray = [...][5]string{
		{"196", "Klaten WareHouse Food", "1000", "Makanan", "Klaten Selatan, Klaten Kab. Jawa Tengah"},
		{"160", "Jember WareHouse", "1000", "Minuman", "Jl. Otto Iskandardinata, Karang Miuwo, Mangli, Kec. Kaliwates, Kabupaten Jember, Jawa Timur 68131"},
		{"133", "Gresik WareHouse Industrial", "1200", "Industrial", "Randuagung, Kec. Kebomas, Kabupaten Gresik, Jawa Timur 61121"},
		{"23", "Bandung WareHouse Komestik", "2000", "Kosmetik", "Jl. Buah Batu No.165, Turangga, Kec. Lengkong, Kota Bandung, Jawa Barat 40264"},
		{"23", "Fashion Bandung WareHouse", "1500", "Fashion", " Jl. Soekarno-Hatta No.825, Babakan Penghulu, Cinambo, Kota Bandung, Jawa Barat 40293"},
		{"21", "Bandar Mataram Lampung Food WareHouse", "1600", "Makanan", "Jalan Soekarno - Hatta, Tanjung Karang Timur, Campang Raya, Bandar Lampung, Kota Bandar Lampung, Lampung 35122"},
		{"318", "Padang WareHouse Liquid", "1700", "Minuman", "Jl. Bagindo Aziz Chan, Batipuh Panjang, Kec. Koto Tangah, Kota Padang, Sumatera Barat 25586"},
		{"289", "Mojokerto Industrial WareHouse", "1800", "Industrial", "Mergelo, Gn. Gedangan, Kec. Magersari, Kota Mojokerto, Jawa Timur 61315"},
		{"285", "Minahasa Kosmetik WareHouse", "1900", "Kosmetik", "JL. Raya Manado Bitung, Manado, 95237, Watutumou, Kalawat, North Minahasa Regency, North Sulawesi"},
		{"278", "Medan Fashion WareHouse", "1200", "Fashion", "Jl. Kayu Putih No.15 A, Tj. Mulia, Kec. Medan Deli, Kota Medan, Sumatera Utara 20241"},
		{"255", "Malang Warehouse Food", "800", "Makanan", "Jl. Letjen Sutoyo No.82, Bunulrejo, Kec. Blimbing, Kota Malang, Jawa Timur 65141"},
	}

	var warehouse entity.Warehouses
	for _, v := range SeedWareHouseArray {
		District, _ := strconv.ParseUint(v[0], 10, 32)
		capacity, _ := strconv.Atoi(v[2])
		warehouse.DistrictsID = uint(District)
		warehouse.WarehousesName = v[1]
		warehouse.WarehousesCapacity = int(capacity)
		warehouse.WarehousesType = v[3]
		warehouse.WarehousesLocation = v[4]
		warehouse.ID = 0
		db.Create(&warehouse)

	}
	fmt.Println("Seeder WareHouse created Sucessfully")
}
