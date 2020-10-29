package user

import "github.com/kostikans/bdProject/models"

type Repository interface {
	CreateNewUser(user models.User) ([]models.User, error)
	GetUserInfo(nickname string) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
}
