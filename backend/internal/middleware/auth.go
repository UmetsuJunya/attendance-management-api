package middleware

import (
	"net/http"
	"strconv"

	UserRepository "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/repositories"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func IsAuth() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Authentication failed.",
				"error":   "Authentication Error",
				"details": "The password provided is incorrect.",
			})
			return
		}
		// before request

		c.Next()
	}
}
