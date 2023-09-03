package view

import (
	"github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/helpers"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func WithGlobalData(c *gin.Context, data gin.H) gin.H {
	data["APP_NAME"] = viper.Get("App.Name")
	data["AUTH"] = helpers.Auth(c)
	return data
}
