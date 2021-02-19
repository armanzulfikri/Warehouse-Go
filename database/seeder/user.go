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
		{"admin", "M", "20/12/1997", "yogyakarta", "admin", "admin@xapiens.id", "admin"},
		{"supervisor", "M", "20/12/1995", "yogyakarta", "supervisor_entry", "supervisor@xapiens.id", "supervisor"},
		{"entry", "M", "20/12/1996", "yogyakarta", "entry", "entry@xapiens.id", "entry"},
		{"manajemen", "M", "20/12/1995", "yogyakarta", "manajemen", "manajemen@xapiens.id", "manajemen"},
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
