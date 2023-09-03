package responses

import (
	userModel "github.com/UmetsuJunya/attendance-management-api/backend/internal/modules/user/models"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

type Users struct {
	Data []User
}

func ToUser(user userModel.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
