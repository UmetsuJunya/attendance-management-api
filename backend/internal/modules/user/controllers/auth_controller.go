package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/requests/auth"
	UserService "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/services"
	"github.com/UmetsuJunya/attendance-management-api/backend/pkg/sessions"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	userService UserService.UserServiceInterface
}

func New() *Controller {
	return &Controller{
		userService: UserService.New(),
	}
}

func (controller *Controller) HandleRegister(c *gin.Context) {
	// validate the request
	var request auth.RegisterRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request. The request data is invalid.",
			"error":   "Invalid request data",
			"details": "The request data is missing required fields or contains invalid data.",
		})
		return
	}

	if controller.userService.CheckUserExists(request.Email) {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User is already registered.",
			"error":   "Duplicate User Registration",
			"details": "The specified email address is already registered.",
		})
		return
	}

	// Create the user
	user, err := controller.userService.Create(request)

	// Check if there is any error on the user creation
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "An error occurred during user registration.",
			"error":   "User Registration Error",
		})
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("The user created successfully with a name %s \n", user.Name)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registration was successful.",
	})
}

func (controller *Controller) HandleLogin(c *gin.Context) {
	// validate the request
	var request auth.LoginRequest
	// This will infer what binder to use depending on the content-type header.
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request. The request data is invalid.",
			"error":   "Invalid request data",
			"details": "The request data is missing required fields or contains invalid data.",
		})
		return
	}

	user, err := controller.userService.HandleUserLogin(request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authentication failed.",
			"error":   "Authentication Error",
			"details": "The password provided is incorrect.",
		})
		return
	}

	sessions.Set(c, "auth", strconv.Itoa(int(user.ID)))

	log.Printf("The user logged in successfully with a name %s \n", user.Name)
	c.JSON(http.StatusAccepted, gin.H{
		"message": "Login successful.",
	})
}

func (controller *Controller) HandleLogout(c *gin.Context) {
	sessions.Remove(c, "auth")

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout was successful.",
	})
}
