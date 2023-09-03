package bootstrap

import (
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/config"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/database"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/html"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/routing"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/sessions"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/static"
)

func Serve() {
	config.Set()

	database.Connect()

	routing.Init()

	sessions.Start(routing.GetRouter())

	static.LoadStatic(routing.GetRouter())

	html.LoadHTML(routing.GetRouter())

	routing.RegisterRoutes()

	routing.Serve()
}
