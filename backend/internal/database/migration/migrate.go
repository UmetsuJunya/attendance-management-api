package migration

import (
	"fmt"
	"log"

	userModels "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/models"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/database"
)

func Migrate() {
	db := database.Connection()

	err := db.AutoMigrate(&userModels.User{})

	if err != nil {
		log.Fatal("Cant migrate")
		return
	}

	fmt.Println("Migration done ..")
}
