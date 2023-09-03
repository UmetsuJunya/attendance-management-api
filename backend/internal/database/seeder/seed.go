package seeder

import (
	"fmt"
	"log"

	userModel "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/models"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/database"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	for i := 1; i <= 10; i++ {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("secret-%d", i)), 12)

		if err != nil {
			log.Fatal("hash password error")
			return
		}

		user := &userModel.User{Name: fmt.Sprintf("Random name %d", i), Email: fmt.Sprintf("random_%d@email.com", i), Password: string(hashedPassword)}

		db.Create(&user) // pass pointer of data to Create

		log.Printf("User created successfully with email address! Name: %s Email: %s \n", user.Name, user.Email)
	}

	log.Println("Seeder done ..")
}
