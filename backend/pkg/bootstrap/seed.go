package bootstrap

import (
	"github.com/UmetsuJunya/attendance-management-api/backend/internal/database/seeder"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/config"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/database"
)

func Seed() {
	config.Set()

	database.Connect()

	seeder.Seed()
}
