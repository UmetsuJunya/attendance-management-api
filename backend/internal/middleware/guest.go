package middleware

import (
	"net/http"
	"strconv"

	UserRepository "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/repositories"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/sessions"
	"github.com/gin-gonic/gin"
)

func IsGuest() gin.HandlerFunc {
	var userRepo = UserRepository.New()

	return func(c *gin.Context) {
		authID := sessions.Get(c, "auth")
		userID, _ := strconv.Atoi(authID)

		user := userRepo.FindByID(userID)

		if user.ID != 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid request. The request data is invalid.",
				"error":   "Invalid request data",
			})
			return
		}
		// before request

		c.Next()
	}
}
