package userUsecase

import (
	"github.com/kostikans/bdProject/models"
	"github.com/kostikans/bdProject/restapi/user"
)

type UserUseCase struct {
	repository user.Repository
}

func NewUserUseCase(repo user.Repository) *UserUseCase {
	return &UserUseCase{repo}
}

func (u *UserUseCase) CreateNewUser(user models.User) ([]models.User, error) {
	return u.repository.CreateNewUser(user)
}

func (u *UserUseCase) GetUserInfo(nickname string) (models.User, error) {
	return u.repository.GetUserInfo(nickname)
}
func (u *UserUseCase) UpdateUser(user models.User) (models.User, error) {
	return u.repository.UpdateUser(user)
}
