package seeder

import (
	"fmt"
	"log"
	"warehouse/entity"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// SeedUser func
func SeedUser(db *gorm.DB) {
	var userArray = [...][7]string{
		{"Arman Zulfikri", "M", "1997-12-20", "yogyakarta", "admin", "arman@xapiens.id", "arman"},
		{"Agus Saputra", "M", "1995-12-20", "yogyakarta", "supervisor_entry", "agus@xapiens.id", "agus"},
		{"Bagus Purnomo", "M", "1996-12-20", "yogyakarta", "entry", "bagus@xapiens.id", "bagus"},
		{"Gery Raharjo", "M", "1995-12-20", "yogyakarta", "manajemen", "gery@xapiens.id", "gery"},
	}

	var user entity.Users
	for _, v := range userArray {
		user.FullName = v[0]
		user.Gender = v[1]
		user.BirthDate = v[2]
		user.BirthPlace = v[3]
		user.Role = v[4]
		user.Email = v[5]
		user.Password = v[6]
		user.ID = 0

		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Error -> ", err.Error())
		}
		user.Password = string(hash)

		db.Create(&user)

	}
	fmt.Println("Seeder Users created Sucessfully")
}
