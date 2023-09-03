package routes

import (
	"github.com/UmetsuJunya/attendance-management-api/backend/internal/middleware"
	userCtrl "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	userController := userCtrl.New()

	guestGroup := router.Group("/")
	guestGroup.Use(middleware.IsGuest())
	{
		guestGroup.POST("/register", userController.HandleRegister)

		guestGroup.POST("/login", userController.HandleLogin)
	}

	authGroup := router.Group("/")
	authGroup.Use(middleware.IsAuth())
	{
		authGroup.POST("/logout", userController.HandleLogout)
	}
}
