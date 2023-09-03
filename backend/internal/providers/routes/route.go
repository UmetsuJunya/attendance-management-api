package routes

import (
	userRoutes "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// ここにルーティングを追加していく
	userRoutes.Routes(router)

}
