package bootstrap

import (
	"github.com/UmetsuJunya/attendance-management-api/backend/internal/database/migration"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/config"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/database"
)

func Migrate() {
	config.Set()

	database.Connect()

	migration.Migrate()
}
