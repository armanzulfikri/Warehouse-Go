package seeder

import (
	"fmt"
	"warehouse/entity"

	"gorm.io/gorm"
)

// SeedSupplier func
func SeedSupplier(db *gorm.DB) {
	var supplierArray = [...][3]string{
		{"Danone Indonesia", "Cyber 2 Tower Lt.12, Jl. H. R. Rasuna Said No.13, RT.7/RW.2, Kuningan, East Kuningan, South Jakarta City, Jakarta 12950", "(021) 29961000"},
		{"PT.Indofood CBP Sukses Makmur Tbk", "Jl. Ipda Tut Harsono No.50-52, Muja Muju, Kec. Umbulharjo, Kota Yogyakarta, Daerah Istimewa Yogyakarta 55165", "(021) 56361000"},
		{"PT. Mayora Indah Tbk", "Ngaliyan, Kec. Ngaliyan, Kota Semarang, Jawa Tengah 50181", "(024) 7626804"},
		{"PT. Sayap Mas Utama Depo Sumedang (WINGS)", "Jl. Parigi Lama No.6, Situ, Kec. Sumedang Utara, Kabupaten Sumedang, Jawa Barat 45621", "(026) 1208391"},
		{"PT. Gudang Garam Tbk.", "Jl. Palagan Tentara Pelajar No.25, Mudal, Sariharjo, Kec. Ngaglik, Kabupaten Sleman, Daerah Istimewa Yogyakarta 55581", "(021) 29961000"},
	}

	var supplier entity.Suppliers
	for _, v := range supplierArray {
		supplier.SupplierName = v[0]
		supplier.Address = v[1]
		supplier.Telepon = v[2]

		supplier.ID = 0

		db.Create(&supplier)

	}
	fmt.Println("Seeder Suppliers created Sucessfully")
}
